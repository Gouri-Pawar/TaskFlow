package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"taskflow/config"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// reads authorization (token)
		authHeader := r.Header.Get("Authorization")

		// no token then error message
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// remove bearer - 'Bearer 'abc.xyz.pqr
		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		// check blacklist first — if this token was logged out, reject immediately
		blacklisted := config.RedisClient.Exists(
			r.Context(),
			"blacklist:"+tokenString,
		).Val()

		if blacklisted > 0 {
			http.Error(
				w,
				"Token has been revoked",
				http.StatusUnauthorized,
			)
			return
		}

		//parse token - header, payload and signature is verified
		//os.Getenv("JWT_SECRET") - used to verify signature

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(
					os.Getenv("JWT_SECRET"),
				), nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(
				w,
				"Invalid Token",
				http.StatusUnauthorized,
			)
			return
		}

		//extract claims (which include data - id, email, etc)
		claims := token.Claims.(jwt.MapClaims)

		//extract user id
		userID := claims["user_id"]

		//every handler can know who is making request
		ctx := context.WithValue(
			r.Context(),
			"user_id",
			userID,
		)

		// pass the raw token too, so the logout handler can blacklist it
		ctx = context.WithValue(
			ctx,
			"token",
			tokenString,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)

	}
}