package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var LecturerInfoCollection *mongo.Collection
var Client *mongo.Client
var err error

type Lecturer struct {
	Email      string `bson:"email" json:"email"`
	Verified   bool   `bson:"verified" json:"verified"`
	FullName   string `bson:"fullName" json:"fullName"`
	GivenName  string `bson:"givenName" json:"givenName"`
	FamilyName string `bson:"familyName" json:"familyName"`
	Picture    string `bson:"picture" json:"picture"`
}

func ConnectDB() error {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://"+os.Getenv("MONGO_USERNAME")+":"+os.Getenv("MONGO_PASSWORD")+"@my-personal-professor-v.k20xc.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		return err
	}
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}
	LecturerInfoCollection = Client.Database("lecturer").Collection("lecturer-info")
	return nil
}

func FindLecturer(email string) (Lecturer, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var lecturer Lecturer
	err := LecturerInfoCollection.FindOne(ctx, bson.D{}).Decode(&lecturer)
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
