package config

import "github.com/caarlos0/env/v10"

var config Config

type Config struct {
	// Application
	Port int `env:"PORT" envDefault:"8082"`

	EmailAPIKey                 string `env:"EMAIL_API_KEY"`
	TaskAssignedEmailTemplateID int64  `env:"TASK_ASSIGNED_EMAIL_TEMPLATE_ID" envDefault:"1"`
}

func LoadEnv() Config {
	return config
}

func SetEnv() error {
	return env.Parse(&config)
}
