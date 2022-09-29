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

	return ctx.NoContent(http.StatusCreated)
}

func (s *secretHandler) GetSecretById(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		log.Error(lib.ErrEmptyParameter)
		return echo.NewHTTPError(http.StatusBadRequest, lib.ErrEmptyParameter)
	}

	secret, err := s.secretService.GetSecret(id)
	if err != nil {
		log.Error("secretService.GetSecret: ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, secret)
}
