package models


type Auth struct {
	Pass						*string							`json:"pass" binding:"required,min=2"					bson:"pass"`
	Email						*string							`json:"email" binding:"required,min=2"				bson:"email"`
}