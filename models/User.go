package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID		 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserCode string 			`json:"user_code,omitempty" bson:"user_code,omitempty"`
	UserName string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Skins	 []int              `json:"skins,omitempty bson:"skins,omitempty`
}