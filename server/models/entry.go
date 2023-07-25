package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// since GOlang doesn't understand JSON, what we are doing is that
// ID is what golang undrstands and "id" is the json object Key.
type Entry struct{
	ID 		primitive.ObjectID  `bson:"id"`
	Dish	*string 		    `json:"dish"`
	Fat		float64             `bson:"fat"`
	Ingredients *string         `json:"ingredients"`
	Calories *string            `json:"calories"`
}