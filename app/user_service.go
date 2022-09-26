package app

import (
	"secret-service/domain"
	"secret-service/dto"
	"secret-service/infrastructure/repository"
	"time"
)

type userService struct {
	repoUser repository.UserRepo
}

func NewUserService(repoUser repository.UserRepo) *userService {
	return &userService{repoUser: repoUser}
}

func (u *userService) Create(req dto.UserCreateReq) (index int, err error) {
	user := domain.User{
		Id:        "",
		Email:     req.Email,
		Password:  req.Password,
		FullName:  req.FullName,
		ShowCount: 0,
		Secret:    "",
		CreatedAt: time.Now(),
	}
	return u.repoUser.CreateUser(user)
}
