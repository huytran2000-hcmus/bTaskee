package config

import "github.com/caarlos0/env/v10"

var config Config

type Config struct {
	// Application
	Port int `env:"PORT" envDefault:"8081"`

	// DB
	DBURI string `env:"DB_URI" envDefault:"mongodb://btaskee:secret@127.0.0.1:27017"`
	// DBHost string `env:"DB_HOST" envDefault:"127.0.0.1"`
	// DBPort string `env:"DB_PORT" envDefault:"27017"`
	// DBUser string `env:"DB_USER" envDefault:"btaskee"`
	// DBPass string `env:"DB_PASS" envDefault:"secret"`
	DBName string `env:"DB_NAME" envDefault:"booking"`
}

func LoadEnv() Config {
	return config
}

func SetEnv() error {
	return env.Parse(&config)
}
