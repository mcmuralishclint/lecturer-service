package models

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type Lecturer struct {
	Email      string `bson:"email" json:"email"`
	Verified   bool   `bson:"verified" json:"verified"`
	FullName   string `bson:"fullName" json:"fullName"`
	GivenName  string `bson:"givenName" json:"givenName"`
	FamilyName string `bson:"familyName" json:"familyName"`
	Picture    string `bson:"picture" json:"picture"`
}

func FindLecturer(email string) (Lecturer, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var lecturer Lecturer
	err := LecturerInfoCollection.FindOne(ctx, bson.M{"email": email}).Decode(&lecturer)
	if err != nil {
		fmt.Println("Error when searching for email: ", err, email)
		return Lecturer{}, nil
	}
	fmt.Println(lecturer)
	return lecturer, nil
}

func CreateLecturer(lecturer Lecturer) (Lecturer, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, insertErr := LecturerInfoCollection.InsertOne(ctx, lecturer)
	if insertErr != nil {
		fmt.Println("InsertOne Error: ", insertErr)
	} else {
		fmt.Println("InsertOne result type: ", reflect.TypeOf(result))
		return lecturer, nil
	}
	return Lecturer{}, nil
}
