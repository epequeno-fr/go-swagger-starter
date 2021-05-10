package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

func init() {
	if err := env.Parse(&Env); err != nil {
		logrus.WithError(err).Fatal("Failed to parse config")
	}

	setupLogrus()
}

var Logger *logrus.Logger

func setupLogrus() {
	Logger = logrus.New()
}
