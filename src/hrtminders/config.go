package main

import (
	"log"

	"github.com/caarlos0/env"
)

type environmentVariables struct {
	BotToken         string `env:"BOT_TOKEN,required"`
	LogLevel         string `env:"LOG_LEVEL" envDefault:"info"`
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort     string `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"legbot"`
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	CronApiUrl       string `env:"CRON_API_URL" envDefault:"https://natural-cron-api.fly.dev"`
}

var ENVIRONMENT *environmentVariables

func initEnv() {
	envVars := &environmentVariables{}
	err := env.Parse(envVars)
	if err != nil {
		log.Fatalf("Error parsing environment variables: %v", err)
	}
	ENVIRONMENT = envVars
}
