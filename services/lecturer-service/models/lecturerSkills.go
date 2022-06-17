package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LecturerSkill struct {
	Email string `bson:"email" json:"email"`
	Skill string `bson:"skill" json:"skill"`
}

func AllLecturerSkills(email string) []string {
	var lecturerSkill LecturerSkill
	cursor, err := LecturerSkillCollection.Find(context.TODO(), bson.D{{"email", email}})
	if err != nil {
		fmt.Println("Error when listing lecturskills [1]: ", err)
		return nil
	}

	var allSkills []string
	for cursor.Next(context.Background()) {
		if err := cursor.Decode(&lecturerSkill); err != nil {
			fmt.Println("Error when listing lecturskills [2]: ", err)
			return nil
		}
		allSkills = append(allSkills, lecturerSkill.Skill)
	}

	return allSkills
}

func AddLecturerSkills(LecturerSkill LecturerSkill) (bool, error) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	filter := bson.D{{"email", LecturerSkill.Email}, {"skill", LecturerSkill.Skill}}
	opts := options.Update().SetUpsert(true)
	_, insertErr := LecturerSkillCollection.UpdateOne(ctx, filter, bson.D{{"$set", bson.D{{"email", LecturerSkill.Email}}}}, opts)
	if insertErr != nil {
		fmt.Println("Error when upserting lectureSkill", insertErr)
		return false, insertErr
	}
	return true, nil
}

func DeleteLecturerSkill(skill string, email string) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := LecturerSkillCollection.DeleteOne(ctx, bson.M{"skill": skill, "email": email})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return nil
}

func DeleteLecturerSkillsBySkill(name_map string) error {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	result, err := LecturerSkillCollection.DeleteMany(ctx, bson.M{"skill": name_map})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	return nil
}
