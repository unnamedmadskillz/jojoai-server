package repository

import (
	"context"
	"test-gpt/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GptRepo struct {
	db *mongo.Collection
}

func NewGptRepo(db *mongo.Database) *GptRepo {
	return &GptRepo{db: db.Collection("suggested")}
}

func (gr *GptRepo) GetAndThrow() model.Suggested {
	var m *model.Suggested
	ctx := context.Background()
	result := gr.db.FindOneAndDelete(ctx, bson.M{"theme": "Окуясу бросил курить", "weight": "5"})
	result.Decode(&m)
	return *m
}
