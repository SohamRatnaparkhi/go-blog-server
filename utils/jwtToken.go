package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Email    string
	Password string
}

type Claims struct {
	creds Credentials
	jwt.RegisteredClaims
}

func GetJwt(signerClaims Credentials) (tokenString string, expireTime time.Time, err error) {
	key := os.Getenv("JWT_SECRET_KEY")
	expiryTime := time.Now().Add(5 * time.Minute)
	claims := Claims{
		creds: signerClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString([]byte(key))

	return tokenString, expireTime, err
}
