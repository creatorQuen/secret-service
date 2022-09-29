package app

import (
	base64 "encoding/base64"
	"github.com/labstack/gommon/log"
	"secret-service/domain"
	"secret-service/dto"
	"secret-service/infrastructure/repository"
)

type secretService struct {
	repoUser repository.UserRepo
}

func NewSecretService(repoUser repository.UserRepo) *secretService {
	return &secretService{repoUser: repoUser}
}

func (s *secretService) PutSecret(req dto.SecretPutReq) error {
	var capsule domain.SecretCapsule
	capsule.Id = req.Id
	secretHash := base64.StdEncoding.EncodeToString([]byte(req.Secret))
	capsule.Secret = secretHash

	err := s.repoUser.InsertSecretById(capsule)
	if err != nil {
		log.Error("repoUser.InsertSecretById: ", err)
		return err
	}

	return nil
}
