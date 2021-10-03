package main

import (
	"context"
	"os"

	"github.com/lab1/errors-logs/api"
	"github.com/lab1/errors-logs/pkg/lib/env"
	"github.com/sirupsen/logrus"
)

func initLogger() *logrus.Logger {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(&logrus.JSONFormatter{})
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	l.SetLevel(logLevel)

	return l
}

func main() {
	l := initLogger()

	apiPort := env.GetIntOrDefault("API_PORT", 5050)

	a := api.Init(l.WithField("logger", "api"), apiPort)

	a.Listen(context.Background())
}
