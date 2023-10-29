package models

type Address struct{
	Id int `json:"id" bson:"_id"`
	Street string `json:"street"  bson:"street"`
	City string `json:"city" bson:"city"`
	State string `json:"state" bson:"state"`
	Zipcode string `json:"zipcode" bson:"zipcode"`
}

type User struct{
	Id int `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password"`
	Address Address `json:"address" bson:"address"`
}