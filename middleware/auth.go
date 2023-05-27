package middleware

import (
	"net/http"
	"os"

	"github.com/SohamRatnaparkhi/go-blog-server/internal/database"
	"github.com/SohamRatnaparkhi/go-blog-server/utils"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey string = os.Getenv("JWT_SECRET_KEY")

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func Auth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		jwtToken, err := req.Cookie("auth-token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tknStr := jwtToken.Value
		claims := &utils.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}
