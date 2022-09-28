package app

import (
	"errors"
	"github.com/labstack/gommon/log"
	"secret-service/domain"
	"secret-service/dto"
	"secret-service/infrastructure/repository"
	"strings"
	"time"
)

type userService struct {
	repoUser repository.UserRepo
}

func NewUserService(repoUser repository.UserRepo) *userService {
	return &userService{repoUser: repoUser}
}

func (u *userService) Create(req dto.UserCreateReq) (index int, err error) {
	var user domain.User
	user.Email = strings.TrimSpace(strings.ToLower(req.Email))

	dbUser, err := u.repoUser.GetUserByEmail(user.Email)
	if err != nil {
		log.Error("repoUser.GetUserByEmail: ", err)
		return 0, err
	}
	if dbUser != nil {
		log.Error("epoUser.GetUserByEmail: ", errors.New("user already exist"))
		return -1, errors.New("user already exist")
	}

	user.FullName = req.FullName
	//user.CreatedAt = time.Now().Format("2006-01-02 15:05:06")
	user.CreatedAt = time.Now()

	var hasher hasherType
	pwdHashed, err := hasher.HashPassword(req.Password)
	if err != nil {
		log.Error("serviceHash.HashPassword: ", err)
		return 0, err
	}
	user.Password = pwdHashed

	return u.repoUser.CreateUser(user)
}
