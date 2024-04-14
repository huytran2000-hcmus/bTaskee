package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/huytran2000-hcmus/bTaskee/booking/config"
	_ "github.com/huytran2000-hcmus/bTaskee/booking/docs"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/infra/pricing_api"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/infra/send_api"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/restful"
)

//	@title			Booking API
//	@version		1.0
//	@description	This is booking apis

// @host	localhost:8080
func main() {
	err := config.SetEnv()
	if err != nil {
		log.Fatalln(err)
	}

	pricingRepo := pricing_api.New()
	sendRepo := send_api.New()
	app, err := restful.New(pricingRepo, sendRepo)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.LoadEnv().Port),
		Handler:      app.Routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  1 * time.Minute,
	}

	err = app.Run(srv)
	if err != nil {
		log.Fatalln(err)
	}
}
