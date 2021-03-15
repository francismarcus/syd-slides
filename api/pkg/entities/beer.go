package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

// Beer entity
type Beer struct {
	ID   primitive.ObjectID `json:"id" bson:"_id, omitemoty"`
	Name string             `json:"name" bson:"name"`
}
