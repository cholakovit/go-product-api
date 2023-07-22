package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID 						primitive.ObjectID 	`bson:"_id,omitempty"`
	Title					string							`json:"title" binding:"required,min=2" 			bson:"title"`
	Desc					string							`json:"desc" 	binding:"required,min=20"			bson:"desc"`
	Tags					string							`json:"tags"	 															bson:"tags"`
	Category_id		string							`json:"category_id"													bson:"category_id"`									
	User_id				string							`json:"user_id"															bson:"user_id"`
	Created_at		time.Time						`json:"created_at"													bson:"created_at"`
	Updated_at		time.Time						`json:"updated_at"													bson:"updated_at"`
}