package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mcmuralishclint/personal_tutor/services/payment-service/handlers"
	"github.com/mcmuralishclint/personal_tutor/services/payment-service/handlers/stripe_handler"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/ping", handlers.Test)

	stripeRoutes := r.Group("/v1/api")
	{
		stripeRoutes.POST("/charges", stripe_handler.ChargeAPI)
	}
	r.Run()
}
