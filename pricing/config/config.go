package config

import "github.com/caarlos0/env/v10"

var config Config

type Config struct {
	// Application
	Port int `env:"PORT" envDefault:"8081"`
}

func LoadEnv() Config {
	return config
}

func SetEnv() error {
	return env.Parse(&config)
}
