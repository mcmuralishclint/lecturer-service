package mongo

import (
	"context"
	"time"

	"github.com/mcmuralishclint/personal_tutor/services/lecturer-service/domain"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
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
	repo := &MongoRepository{
		client:  mongoClient,
		db:      mongoDb,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, errors.Wrap(err, "client error")
	}

	return repo, nil
}
