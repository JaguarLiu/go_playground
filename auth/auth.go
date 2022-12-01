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
	VerifyToken(token string)
}
type Auth struct{}

func (srv *Auth) Login(user User) error {
	if user.Account != "admin" && user.Password != "1234" {
		return errors.New("password or account incorrect")
	}
	return nil
}
func (srv *Auth) CreateToken(user User) (string, error) {
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
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString("123456")
	if err != nil {
		return "", err
	}
	return token, nil
}

func (srv *Auth) VerifyToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("123456"), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
