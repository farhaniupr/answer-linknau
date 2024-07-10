package main

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const JWTKey = "abcdefghijkl"

func ValidasiToken(tokenString string) (map[string]interface{}, bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})

	if token.Valid {
		return token.Claims.(jwt.MapClaims), true, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, false, errors.New("token malformed")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, false, errors.New("token invalid")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return nil, false, errors.New("token expired")
	}

	return nil, false, errors.New("couldn't handle token")
}

func CreateToken(dataReq User) (string, error) {
	mySigningKey := []byte(JWTKey)

	type MyCustomClaims struct {
		Phone string `json:"phone"`
		Name  string `json:"name"`
		jwt.RegisteredClaims
	}

	claims := MyCustomClaims{
		dataReq.Phone,
		dataReq.Name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "https://localhost:8080",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenGenerate, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenGenerate, nil
}
