package config

import "github.com/caarlos0/env/v10"

var config Config

type Config struct {
	// Application
	Port int `env:"PORT" envDefault:"8080"`

	// DB
	DBURI string `env:"DB_URI" envDefault:"mongodb://btaskee:secret@127.0.0.1:27017"`
	// DBHost string `env:"DB_HOST" envDefault:"127.0.0.1"`
	// DBPort string `env:"DB_PORT" envDefault:"27017"`
	// DBUser string `env:"DB_USER" envDefault:"btaskee"`
	// DBPass string `env:"DB_PASS" envDefault:"secret"`
	DBName string `env:"DB_NAME" envDefault:"booking"`

	// Pricing API
	CalcalatePriceURI string `env:"CALCULATE_PRICE_URI" envDefault:"http://localhost:8081/api/v1/task:calculate-price"`

	// Send API
	SendTaskURI string `env:"SEND_TASK_ID" envDefault:"http://localhost:8082/api/v1/task:send"`
}

func LoadEnv() Config {
	return config
}

func SetEnv() error {
	return env.Parse(&config)
}
