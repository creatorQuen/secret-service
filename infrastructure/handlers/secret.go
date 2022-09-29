package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"secret-service/app"
	"secret-service/dto"
	"secret-service/lib"
)

type secretHandler struct {
	secretService app.SecretService
}

func NewSecretHandler(secretService app.SecretService) *secretHandler {
	return &secretHandler{secretService: secretService}
}

func (s *secretHandler) CreateSecret(ctx echo.Context) error {
	var req dto.SecretPutReq
	err := ctx.Bind(&req)
	if err != nil {
		log.Error("Request to bind ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if req.Secret == "" {
		log.Error(lib.ErrSecretIsEmpty)
		return echo.NewHTTPError(http.StatusBadRequest, lib.ErrSecretIsEmpty)
	}

	err = s.secretService.PutSecret(req)
	if err != nil {
		log.Error("secretService.PutSecret: ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return echo.NewHTTPError(http.StatusCreated)
}
