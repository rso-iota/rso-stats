package config

import (
	"reflect"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	NatsURL string `env:"NATS_URL"`
	LogJSON bool   `env:"LOG_JSON"`
}

func Init() Config {
	godotenv.Load("defaults.env")

	var config Config
	err := env.Parse(&config)
	if err != nil {
		log.WithError(err).Fatal("Failed to parse config")
	}

	fields := log.Fields{}

	val := reflect.ValueOf(config)
	for i := 0; i < val.NumField(); i++ {
		fields[val.Type().Field(i).Name] = val.Field(i).Interface()
	}

	log.WithFields(fields).Info("Loaded config")

	if config.LogJSON {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetLevel(log.DebugLevel)

	return config
}
