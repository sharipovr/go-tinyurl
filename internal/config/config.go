package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"PORT" default:"8080"`
}

func MustLoad() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		log.Fatalf("load config: %v", err)
	}
	return &c
}
