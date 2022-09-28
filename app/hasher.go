package app

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type hasherType struct{}

func (h hasherType) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error generating salt and hashing password using bcrypt: %w", err)
	}

	return string(hash), err
}

func (h hasherType) IsPasswordCorrect(providedPassword string, storedHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedPassword)) == nil
}
