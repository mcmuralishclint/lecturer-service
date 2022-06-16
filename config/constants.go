package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func InitMongo() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://"+os.Getenv("MONGO_USERNAME")+":"+os.Getenv("MONGO_PASSWORD")+"@my-personal-professor-v.k20xc.mongodb.net/?retryWrites=true&w=majority"))
	return client, err
}
