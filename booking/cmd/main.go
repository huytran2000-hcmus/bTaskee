package main

import (
	"log"

	"github.com/huytran2000-hcmus/bTaskee/booking/config"
	_ "github.com/huytran2000-hcmus/bTaskee/booking/docs"
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

	app, err := restful.New()
	if err != nil {
		log.Fatalln(err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
