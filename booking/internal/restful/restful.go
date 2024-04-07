package restful

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"log"

	"github.com/huytran2000-hcmus/bTaskee/booking/config"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/infra/mongodb"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/infra/pricing_api"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/infra/send_api"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/service"
	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type restful struct {
	e  *echo.Echo
	wg sync.WaitGroup
}

func New() (*restful, error) {
	e := echo.New()

	// repositories
	mongoDB, err := mongodb.NewRepository(config.LoadEnv().DBName, config.LoadEnv().DBURI)
	if err != nil {
		return nil, err
	}
	pricingRepo := pricing_api.New()
	sendRepo := send_api.New()

	// services
	taskSVC := service.NewTask(mongoDB, pricingRepo, sendRepo)

	// handlers
	taskH := newTaskHandler(taskSVC)

	apiBase := e.Group("/api")
	apiV1 := apiBase.Group("/v1")

	apiTask := apiV1.Group("/tasks")
	{
		apiTask.POST("", taskH.createTaskHandler)
		apiTask.GET("/:id", taskH.getTaskByID)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	return &restful{
		e: e,
	}, nil
}

func (r *restful) Run() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.LoadEnv().Port),
		Handler:      r.e,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	shutDownErr := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit
		log.Printf("shutting down server: %s", s.String())
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			shutDownErr <- err
		}

		log.Println("completing background tasks")

		r.wg.Wait()
		shutDownErr <- nil
	}()

	log.Printf("starting server at port %d", config.LoadEnv().Port)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutDownErr
	if err != nil {
		return err
	}

	log.Println("stopped server")

	return nil
}
