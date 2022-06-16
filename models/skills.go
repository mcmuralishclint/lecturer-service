package models

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Skill struct {
	NameMap string `bson:"name_map" json:"name_map"`
	Value   string `bson:"value" json:"value"`
}

func DeleteSkill(nameMap string) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := SkillsCollection.DeleteOne(ctx, bson.M{"name_map": nameMap})
	if err != nil {
		return err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return nil
}

func FindSkill(nameMap string) (bool, Skill) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var skill Skill
	err := SkillsCollection.FindOne(ctx, bson.M{"name_map": nameMap}).Decode(&skill)
	return err == nil, skill
}

func Skills() *mongo.Cursor {
	cursor, err := SkillsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Error when listing skills: ", err)
		return nil
	}

	return cursor
}

func CreateSkill(skill Skill) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	exists, _ := FindSkill(skill.NameMap)
	if exists {
		return errors.New("item already exists")
	}
	result, insertErr := SkillsCollection.InsertOne(ctx, skill)
	if insertErr != nil {
		fmt.Println("InsertOne Error: ", insertErr)
		return insertErr
	} else {
		fmt.Println("InsertOne result type: ", reflect.TypeOf(result))
		return nil
	}
}
