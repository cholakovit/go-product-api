package models

import (
	 "go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `json:"title" binding:"required,min=2" 			bson:"title"`
	Img          string             `json:"img"																	bson:"img"`
	Desc         string             `json:"desc" binding:"required,min=20"			bson:"desc"`
	Short_Desc   string             `json:"short_desc"													bson:"short_desc"`
	Price        float64            `json:"price" binding:"required"						bson:"price"`
	Discount     float64            `json:"discount"														bson:"discount"`
	Stock        int                `json:"stock"																bson:"stock"`
	Active       bool               `json:"active"															bson:"active"`
	Manufacturer string             `json:"manufacturer"												bson:"manufacturer"`
	Thumbs       Thumbs             `json:"thumbs"															bson:"thumbs"`
	Category_id	 string							`json:"category_id"													bson:"category_id"`
	User_id			 string							`json:"user_id"															bson:"user_id"`
}

type Thumbs struct {
	Thumb1 string `json:"thumb1" bson:"thumb1"`
	Thumb2 string `json:"thumb2" bson:"thumb2"`
	Thumb3 string `json:"thumb3" bson:"thumb3"`
	Thumb4 string `json:"thumb4" bson:"thumb4"`
	Thumb5 string `json:"thumb5" bson:"thumb5"`
}