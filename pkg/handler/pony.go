package handler

import (
	"github.com/go-openapi/runtime/middleware"

	"example.ponies.com/api/swagger_gen/models"
	"example.ponies.com/api/swagger_gen/server/restapi/pony"
)

func (h *handler) GetPony(params pony.GetPonyParams, principal *models.Principal) middleware.Responder {
	h.log.Info("getting a pony")
	return pony.NewGetPonyOK().WithPayload(&models.PonyResponse{
		Color: "blue",
		Name:  "Jane",
	})
}

func (h *handler) PutPony(params pony.UpsertPonyParams, principal *models.Principal) middleware.Responder {
	h.log.Info("putting a pony")
	return pony.NewUpsertPonyOK().WithPayload(&models.PonyResponse{
		Color: *params.Body.Color,
		Name:  *params.Body.Name,
	})
}
