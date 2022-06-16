package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GoogleClientID"),
		ClientSecret: os.Getenv("GoogleClientSecret"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
	return conf
}
