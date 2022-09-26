package app

import "secret-service/dto"

type UserService interface {
	Create(req dto.UserCreateReq) (int, error)
}
