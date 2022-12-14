package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"secret-service/app"
	"secret-service/dto"
)

type userHandler struct {
	userService app.UserService
}

func NewUserHandler(userService app.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (u *userHandler) CreateUser(ctx echo.Context) error {
	var req dto.UserCreateReq
	err := ctx.Bind(&req)
	if err != nil {
		log.Println("Request to bind ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := u.userService.Create(req)
	if err != nil {
		log.Println("userService.Create: ", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusCreated, id)
}
