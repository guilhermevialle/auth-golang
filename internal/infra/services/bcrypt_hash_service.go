package infra_services

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptHashService struct{}

// implements IHashService
var _ IHashService = (*BcryptHashService)(nil)

func NewBcryptHashService() *BcryptHashService {
	return &BcryptHashService{}
}

func (h *BcryptHashService) Hash(plain string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (h *BcryptHashService) Compare(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
