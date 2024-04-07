package main

import (
	"log"

	"github.com/huytran2000-hcmus/bTaskee/send/config"
	"github.com/huytran2000-hcmus/bTaskee/send/internal/restful"
)

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