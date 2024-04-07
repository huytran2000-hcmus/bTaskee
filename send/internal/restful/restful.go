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

	"github.com/huytran2000-hcmus/bTaskee/send/config"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/infra/brevo"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/service"
	echo "github.com/labstack/echo/v4"
)

type restful struct {
	e  *echo.Echo
	wg sync.WaitGroup
}

func New() (*restful, error) {
	e := echo.New()

	// repositories
	emailRepo := brevo.New()

	// services
	sendSVC := service.NewSend(emailRepo)

	// handlers
	sendH := NewSendHandler(sendSVC)

	apiBase := e.Group("/api")
	apiV1 := apiBase.Group("/v1")
	{
		apiV1.POST("/task:send", sendH.send)
	}

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
