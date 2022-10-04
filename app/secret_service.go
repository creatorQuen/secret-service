package app

import (
	base64 "encoding/base64"
	"github.com/labstack/gommon/log"
	"secret-service/domain"
	"secret-service/dto"
	"secret-service/infrastructure/repository"
	"secret-service/lib"
	"unicode/utf8"
)

type secretService struct {
	repoUser repository.UserRepo
}

func NewSecretService(repoUser repository.UserRepo) *secretService {
	return &secretService{repoUser: repoUser}
}

func (s *secretService) PutSecret(req dto.SecretPutReq) error {
	if utf8.RuneCountInString(req.Secret) > lib.MaxLengthSecret {
		return lib.ErrMaxLengthString
	}

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

func (s *secretService) GetSecret(id string) (string, error) {
	showCount, err := s.repoUser.GetShowCountById(id)
	if err != nil {
		return "", err
	}
	if showCount >= lib.ShowCount {
		return "", lib.ErrDontShowCount
	}

	secretHash, err := s.repoUser.SelectSecretAndAddCountById(id)
	if err != nil {
		return "", err
	}

	secret, err := base64.StdEncoding.DecodeString(secretHash)
	if err != nil {
		return "", err
	}

	return string(secret), nil
}
