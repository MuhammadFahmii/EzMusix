package middlewares

import (
	"EzMusix/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserId int
	jwt.StandardClaims
}

func GenerateToken(userId int) (string, error) {
	claims := JwtClaims{
		userId,
		jwt.StandardClaims{ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(constants.SECRET_JWT))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
