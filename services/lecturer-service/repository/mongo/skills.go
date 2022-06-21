package mongo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *MongoRepository) FindAllSkills() []domain.Skill {
	var allSkills []domain.Skill
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.db).Collection("skills")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return make([]domain.Skill, 0)
	}
	skill := domain.Skill{}
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&skill); err != nil {
			log.Fatal(err)
		}
		allSkills = append(allSkills, skill)
	}
	return allSkills
}

func (r *MongoRepository) CreateSkill(skill domain.Skill) (domain.Skill, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	collection := r.client.Database(r.db).Collection("skills")
	result, insertErr := collection.InsertOne(ctx, skill)
	if insertErr != nil {
		fmt.Println("InsertOne Error: ", insertErr)
		return domain.Skill{}, insertErr
	} else {
		fmt.Println("InsertOne result type: ", reflect.TypeOf(result))
		return skill, nil
	}
}
func (r *MongoRepository) FindSkill(name_map string) domain.Skill {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var skill domain.Skill
	collection := r.client.Database(r.db).Collection("skills")
	err := collection.FindOne(ctx, bson.M{"name_map": name_map}).Decode(&skill)
	if err != nil {
		return domain.Skill{}
	}
	return skill
}
func (r *MongoRepository) DeleteSkill(name_map string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	collection := r.client.Database(r.db).Collection("skills")
	result, err := collection.DeleteOne(ctx, bson.M{"name_map": name_map})
	if err != nil {
		return false, err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return true, nil
}
