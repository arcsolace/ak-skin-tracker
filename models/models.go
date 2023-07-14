package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID		    primitive.ObjectID`json:"_id,omitempty" bson:"_id,omitempty"`
	UserCode    string 			   `json:"user_code,omitempty" bson:"user_code,omitempty"`
	UserName    string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Skins	    []int              `json:"skins,omitempty bson:"skins,omitempty`
}

type SkinBrief struct {
	ID		    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SkinId		int				   `json:"skin_id,omitempty" bson:"skin_id,omitempty"`
	Skin		string			   `json:"skin,omitempty" bson:"skin,omitempty"`
	Brand		int				   `json:"brand,omitempty" bson:"brand,omitempty"`
	Rarity		int				   `json:"rarity,omitempty" bson:"rarity,omitempty"`
	CnRelease   string             `json:"cn_release,omitempty" bson:"cn_release,omitempty"`
	Image		string			   `json:"image,omitempty" bson:"image,omitempty"`
	Live2d		bool			   `json:"live2d,omitempty" bson:"live2d,omitempty"`
	Event		bool			   `json:"event,omitempty" bson:"event,omitempty"`
}

type Skin struct {
	ID		    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	SkinId		int				   `json:"skin_id,omitempty" bson:"skin_id,omitempty"`
	Skin		string			   `json:"skin,omitempty" bson:"skin,omitempty"`
	Brand		int				   `json:"brand,omitempty" bson:"brand,omitempty"`
	Cost		string			   `json:"cost,omitempty" bson:"cost,omitempty"`
	Rarity		int				   `json:"rarity,omitempty" bson:"rarity,omitempty"`
	NaRelease   string             `json:"na_release,omitempty" bson:"na_release,omitempty"`
	CnRelease   string             `json:"cn_release,omitempty" bson:"cn_release,omitempty"`
	Image		string			   `json:"image,omitempty" bson:"image,omitempty"`
	LargeImage  string			   `json:"large_image,omitempty" bson:"large_image,omitempty"`
	Live2d		bool			   `json:"live2d,omitempty" bson:"live2d,omitempty"`
	Event		bool			   `json:"event,omitempty" bson:"event,omitempty"`
}

type Brand struct {
	ID		    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	BrandId		int				   `json:"brand_id,omitempty" bson:"brand_id,omitempty"`
	BrandName	string			   `json:"brand_name,omitempty" bson:"brand_name,omitempty"`
	BrandImage	string			   `json:"brand_image,omitempty" bson:"brand_image,omitempty"`
}