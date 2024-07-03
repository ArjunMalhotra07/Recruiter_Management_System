package mymiddleware

import (
	"context"
	"net/http"
	"strings"

	apigateway "github.com/ArjunMalhotra07/Recruiter_Management_System/api_gateway"
	"github.com/ArjunMalhotra07/Recruiter_Management_System/handler"
	"github.com/ArjunMalhotra07/Recruiter_Management_System/models"
	"github.com/dgrijalva/jwt-go"
)

func JwtVerify(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				response := models.Response{Message: "Missing Authorization header", Status: "Error"}
				handler.SendResponse(w, response)
				return
			}

			tokenString := strings.TrimSpace(strings.Replace(authorizationHeader, "Bearer", "", 1))
			token, err := apigateway.VerifyToken(tokenString, secret)
			if err != nil {
				response := models.Response{Message: err.Error(), Status: "Error"}
				handler.SendResponse(w, response)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
				next.ServeHTTP(w, r)
			} else {
				response := models.Response{Message: "Invalid Token", Status: "Error"}
				handler.SendResponse(w, response)
				return
			}
		})
	}
}
