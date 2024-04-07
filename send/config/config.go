package config

import "github.com/caarlos0/env/v10"

var config Config

type Config struct {
	// Application
	Port int `env:"PORT" envDefault:"8082"`

	BrevoAPIKey                      string `env:"BREVO_API_KEY" envDefault:"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"`
	BrevoTaskAssignedEmailTemplateID int64  `env:"TASK_ASSIGNED_EMAIL_TEMPLATE_ID" envDefault:"1"`
}

func LoadEnv() Config {
	return config
}

func SetEnv() error {
	return env.Parse(&config)
}
