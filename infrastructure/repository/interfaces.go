package repository

import "secret-service/domain"

type UserRepo interface {
	CreateUser(user domain.User) (int, error)
}
