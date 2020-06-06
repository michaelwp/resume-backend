package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Biodata struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"first_name"`
	BirthPlace string `json:"birth_place"`
	BirthDate time.Time `json:"birth_place"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Contact struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	BiodataId string `json:"biodata_id"`
}

type SocialMedia struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
	BiodataId string `json:"biodata_id"`
}

type BiodataResponse struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"first_name"`
	BirthPlace string `json:"birth_place"`
	BirthDate time.Time `json:"birth_place"`
	Address string `json:"address"`
	City string `json:"city"`
	Province string `json:"province"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	Contact Contact
	SocialMedia SocialMedia
}

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data BiodataResponse
}