package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	"example.ponies.com/api/pkg/config"
	"example.ponies.com/api/swagger_gen/models"
	"example.ponies.com/api/swagger_gen/server/restapi"
	"example.ponies.com/api/swagger_gen/server/restapi/health"
	"example.ponies.com/api/swagger_gen/server/restapi/pony"
)

type Handlers interface {
	GetPony(params pony.GetPonyParams, principal *models.Principal) middleware.Responder
	PutPony(params pony.UpsertPonyParams, principal *models.Principal) middleware.Responder
}

type handler struct {
	log *logrus.Logger
}

func newHandler() Handlers {
	return &handler{
		log: config.Logger,
	}
}

func Setup(api *restapi.PoniesAPI) {
	h := newHandler()
	api.PonyGetPonyHandler = pony.GetPonyHandlerFunc(h.GetPony)
	api.PonyUpsertPonyHandler = pony.UpsertPonyHandlerFunc(h.PutPony)

	// Health
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(
		func(params health.GetHealthParams) middleware.Responder {
			return health.NewGetHealthOK().WithPayload(&models.Health{Status: "OK"})
		},
	)
}
