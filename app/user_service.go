package app

import (
	"errors"
	"github.com/labstack/gommon/log"
	"secret-service/domain"
	"secret-service/dto"
	"secret-service/infrastructure/repository"
	"secret-service/lib"
	"strings"
	"time"
)

type userService struct {
	repoUser repository.UserRepo
}

func NewUserService(repoUser repository.UserRepo) *userService {
	return &userService{repoUser: repoUser}
}

func (u *userService) Create(req dto.UserCreateReq) (index string, err error) {
	var user domain.User

	user.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if !isValidEmail(req.Email) {
		log.Error(lib.ErrNotValidEmail)
		return "", lib.ErrNotValidEmail
	}

	if !isValidPassword(req.Password) {
		log.Error(lib.ErrNotValidPassword)
		return "", lib.ErrNotValidPassword
	}

	dbUser, err := u.repoUser.GetUserByEmail(user.Email)
	if err != nil {
		log.Error("repoUser.GetUserByEmail: ", err)
		return "", err
	}
	if dbUser != nil {
		log.Error("repoUser.GetUserByEmail: ", errors.New("user already exist"))
		return "", errors.New("user already exist")
	}

	user.FullName = req.FullName
	user.CreatedAt = time.Now().Format(lib.DbTLayout)
	user.Id = lib.GetUUID()

	var hasher hasherType
	pwdHashed, err := hasher.HashPassword(req.Password)
	if err != nil {
		log.Error("serviceHash.HashPassword: ", err)
		return "", err
	}
	user.Password = pwdHashed

	return u.repoUser.CreateUser(user)
}
