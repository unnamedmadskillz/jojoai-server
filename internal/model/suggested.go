package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Suggested struct {
	ID     primitive.ObjectID `bson:"_id"`
	Theme  string             `bson:"theme"`
	Weight int                `bson:"value"`
}
