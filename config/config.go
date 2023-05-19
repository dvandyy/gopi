package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v8"
)

type Config struct {
	Environment string `env:"HOME"`
	Version     string `env:"VERSION"`

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

var Cfg = &Config{}

func New(logger *log.Logger) {
	err := env.Parse(Cfg) // Parsing env values
	if err != nil {
		logger.Printf("%+v\n", err)
	}
}
