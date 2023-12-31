package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReplicRow struct {
	Name  string `bson:"model_name"`
	Text  string `bson:"tts_text"`
	Path  string `bson:"output_file_path"`
	Order int    `bson:"order"`
}

type ReplicDB struct {
	ID        primitive.ObjectID `bson:"_id"`
	Data      []ReplicRow        `bson:"data"`
	CreatedAt time.Time          `bson:"created_at"`
}
