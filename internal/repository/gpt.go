package repository

import (
	"context"
	"fmt"
	"test-gpt/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GptRepo struct {
	db *mongo.Database
}

func NewGptRepo(db *mongo.Database) *GptRepo {
	return &GptRepo{db: db}
}

func (gr *GptRepo) GetSuggestedAndThrow() model.Suggested {
	var m model.Suggested
	ctx := context.Background()
	result := gr.db.Collection(collectionSuggested).FindOneAndDelete(ctx, bson.M{"theme": "Окуясу бросил курить", "weight": "5"})
	result.Decode(&m)
	return m
}

func (gr *GptRepo) PutCompletedDialogue(d *model.ReplicDB) (*mongo.InsertOneResult, error) {
	ctx := context.Background()
	result, err := gr.db.Collection(collectionCreated).InsertOne(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("failed to insert dialogue, result: %v, %v", result, err)
	}
	return result, nil
}
