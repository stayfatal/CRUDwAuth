package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("super_secret_key")

type Claims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func CreateToken(id int) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 72).Unix()
	claims := &Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
