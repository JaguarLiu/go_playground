package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	Account  string
	Password string
}

type AuthSrv interface {
	Login(user User)
	CreateToken(user User)
}
type Auth struct{}

func (srv *Auth) Login(user User) error {
	if user.Account != "admin" && user.Password != "1234" {
		return errors.New("password or account incorrect")
	}
	return nil
}
func (srv *Auth) CreateToken(user User) {
	now := time.Now()
	claims := jwt.StandardClaims{
		Audience:  user.Account,
		ExpiresAt: now.Add(20 * time.Second).Unix(),
		Id:        user.Account,
		IssuedAt:  now.Unix(),
		Issuer:    "demoJWT",
		NotBefore: now.Add(10 * time.Second).Unix(),
		Subject:   user.Account,
	}
}
