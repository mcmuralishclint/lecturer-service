package models

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
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

func AddLecturerSkills() {

}
