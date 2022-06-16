package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/config"
	"github.com/mcmuralishclint/personal_tutor/lecturer-service/models"
)

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		claims, success := extractClaims(token)
		if success {
			if claims["authorized"] == "" || claims["authorized"] == false {
				http.Error(w, "Forbidden", http.StatusForbidden)
			} else {
				next.ServeHTTP(w, r)
			}
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte(config.MySigningKey)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	isAdmin, _ := models.IsAdmin(email)
	claims["authorized"] = isAdmin
	claims["user"] = email
	claims["exp"] = time.Now().Add(time.Minute + 30).Unix()

	tokenString, err := token.SignedString(config.MySigningKey)
	if err != nil {
		fmt.Printf("Something went wrong %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
