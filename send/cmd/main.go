package main

import (
	"log"

	"github.com/huytran2000-hcmus/bTaskee/send/config"
	_ "github.com/huytran2000-hcmus/bTaskee/send/docs"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/restful"
)

//	@title			Send API
//	@version		1.0
//	@description	This is send apis

// @host	localhost:8082
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
