package repository

import (
	"context"
	"log"
	"time"

	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	client  *mongo.Client
	db      string
	timeout time.Duration
}

func newMongoClient(mongoServerURL string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoServerURL))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoServerURL string, mongoDb string, timeout int) (domain.Repository, error) {
	mongoClient, err := newMongoClient(mongoServerURL, timeout)
	repo := &mongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	return repo, nil
}

func (r *mongoRepository) FindAll() []domain.Skill {
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

func (r *mongoRepository) Create() (bool, error) {
	return true, nil
}
func (r *mongoRepository) Find(name_map string) domain.Skill {
	return domain.Skill{}
}
func (r *mongoRepository) Delete() (bool, error) {
	return true, nil
}
