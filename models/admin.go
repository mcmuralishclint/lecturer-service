package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Admin struct {
	Email string `bson:"email" json:"email"`
}

func IsAdmin(email string) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var admin Admin
	err := AdminCollection.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err != nil {
		fmt.Println("Error when searching for admin by email: ", err, email)
		return false, nil
	}
	return true, nil
}
