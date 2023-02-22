package util

import (
	"email_auth/initer"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string) (string, error) {

	tokenStruct := TokenClaims{
		email,
		jwt.RegisteredClaims{
			Issuer:    initer.AppConfig.JwtIssuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(initer.AppConfig.JwtExpHour))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenStruct)

	tokenString, err := token.SignedString([]byte(initer.AppConfig.JwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token string) (*TokenClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(initer.AppConfig.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("sdasdasd")

	if claims, ok := parsedToken.Claims.(*TokenClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
