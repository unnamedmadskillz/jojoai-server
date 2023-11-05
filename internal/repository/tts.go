package repository

import (
	"context"
	"log"
	"test-gpt/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TTSRepo struct {
	db *mongo.Database
}

func NewTTSRepo(db *mongo.Database) *TTSRepo {
	return &TTSRepo{
		db: db,
	}
}

func (ttsrepo *TTSRepo) GetDialogue() *model.ReplicDB {
	var rdb *model.ReplicDB
	ctx := context.Background()
	opts := options.FindOne().SetSort(bson.M{"created_at": 1})
	result := ttsrepo.db.Collection(collectionCreated).FindOne(ctx, bson.D{}, opts)
	err := result.Decode(&rdb)
	if err != nil {
		log.Printf("Failed to get dialogue for tts: %v", err)
	}
	return rdb
}
