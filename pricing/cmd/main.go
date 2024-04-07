package main

import (
	"log"

	"github.com/huytran2000-hcmus/bTaskee/pricing/config"
	_ "github.com/huytran2000-hcmus/bTaskee/pricing/docs"
	"github.com/huytran2000-hcmus/bTaskee/pricing/internal/restful"
)

//	@title			Pricing API
//	@version		1.0
//	@description	This is pricing apis

// @host	localhost:8081
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
