package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 							primitive.ObjectID	`bson:"_id"`
	FullName				*string							`json:"full_name"		bson:"full_name"`
	Pass						*string							`json:"pass"				bson:"pass"`
	Email						*string							`json:"email"				bson:"email"`
	Phone						*string							`json:"phone"				bson:"phone"`
	Role						*string							`json:"role"				bson:"role"`
	Token						*string							`json:"token"				bson:"token"`
	Rtoken					*string							`json:"rtoken"			bson:"rtoken"`
	Created_at			time.Time						`json:"created_at"	bson:"created_at"`
	Updated_at			time.Time						`json:"updated_at"	bson:"updated_at"`
	User_id					string							`json:"user_id"			bson:"user_id"`
}