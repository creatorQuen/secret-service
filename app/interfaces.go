package app

import "secret-service/dto"

type UserService interface {
	Create(req dto.UserCreateReq) (string, error)
}

type Hasher interface {
	HashPassword(password string) (string, error)
}

type SecretService interface {
	PutSecret(req dto.SecretPutReq) error
	GetSecret(id string) (string, error)
}
