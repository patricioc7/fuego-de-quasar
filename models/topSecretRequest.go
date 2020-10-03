package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Satellite struct {
	Name string `json:"name"`
	Distance float32 `json:"distance"`
	Message []string `json:"message"`
}

type TopSecret struct {
	Id     	  primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Satellites  [3]Satellite `json:"satellites"`
}