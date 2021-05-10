package config

import (
	"net/http"

	"example.ponies.com/api/pkg/constants"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	n := negroni.New()

	// Recovery logger to catch panics
	n.Use(setupRecoveryMiddleware())

	n.UseHandler(handler)

	if Env.MiddlewareVerboseLoggerEnabled {
		middleware := negronilogrus.NewMiddlewareFromLogger(Logger, constants.ApiName)
		for _, u := range Env.MiddlewareVerboseLoggerExcludeURLs {
			if err := middleware.ExcludeURL(u); err != nil {
				logrus.WithError(err).Error("Failed to exclude URL from middleware")
			}
		}
		n.Use(middleware)
	}

	return n
}

type recoveryLogger struct{}

func (r *recoveryLogger) Printf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}

func (r *recoveryLogger) Println(v ...interface{}) {
	logrus.Errorln(v...)
}

func setupRecoveryMiddleware() *negroni.Recovery {
	r := negroni.NewRecovery()
	r.Logger = &recoveryLogger{}
	return r
}
