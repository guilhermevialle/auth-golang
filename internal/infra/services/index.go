package infra_services

import "time"

type ITokenService interface {
	Generate(id string, secret string, expTime time.Duration) (string, error)
	Validate(token string, secret string) (string, error)
}

type IHashService interface {
	Hash(plain string) (string, error)
	Compare(hashed, plain string) bool
}
