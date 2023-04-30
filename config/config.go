package config

import (
	"github.com/caarlos0/env/v8"
	"log"
	"time"
)

type Config struct {
	Environment string  `env:"HOME"`
	Version     float32 `env:"VERSION"`

	Address      string        `env:"ADDRESS"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT"`

	DB_Host     string `env:"DB_HOST"`
	DB_Port     int    `env:"DB_PORT"`
	DB_User     string `env:"DB_USER"`
	DB_Password string `env:"DB_PASSWORD"`
	DB_Name     string `env:"DB_NAME"`
}

func New(logger *log.Logger) *Config {
	cfg := &Config{}
	err := env.Parse(cfg) // Parsing env values
	if err != nil {
		logger.Printf("%+v\n", err)
	}
	return cfg
}
