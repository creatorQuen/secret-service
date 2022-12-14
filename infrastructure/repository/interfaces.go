package repository

import (
	"secret-service/domain"
)

type UserRepo interface {
	CreateUser(user domain.User) (string, error)
	GetUserByEmail(string) (*domain.User, error)
	InsertSecretById(domain.SecretCapsule) error
	SelectSecretAndAddCountById(id string) (string, error)
	GetShowCountById(id string) (int, error)
}
