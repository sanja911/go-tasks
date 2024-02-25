package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id ,omitempty" bson:"_id"`
	Name     string             `json:"name,omitempty" validate:"required" bson:"name"`
	Email    string             `json:"email,omitempty" validate:"required" bson:"email"`
	Location string             `json:"location,omitempty" validate:"required" bson:"location"`
	Password string             `json:"password" validate:"required" bson:"password"`
}
