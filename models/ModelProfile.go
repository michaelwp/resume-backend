package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Biodata struct {
	Id 			primitive.ObjectID `bson:"_id" json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	BirthPlace 	string `json:"birth_place"`
	BirthDate 	string `json:"birth_date"`
	Address 	string `json:"address"`
	City 		string `json:"city"`
	Province 	string `json:"province"`
	State 		string `json:"state"`
	ZipCode 	string `json:"zip_code"`
}

type AboutMe struct {
	Id 			primitive.ObjectID `bson:"_id" json:"id"`
	AboutMe		string `json:"about_me"`
	BiodataId 	string `json:"biodata_id"`
}

type Contact struct {
	Id 			primitive.ObjectID `bson:"_id" json:"id"`
	PhoneNumber string `json:"phone_number"`
	Email 		string `json:"email"`
	BiodataId 	string `json:"biodata_id"`
}

type SocialMedia struct {
	Id 			primitive.ObjectID `bson:"_id" json:"id"`
	Name 		string `json:"name"`
	Link 		string `json:"link"`
	BiodataId 	string `json:"biodata_id"`
}

type ProfilePicture struct {
	Id			primitive.ObjectID `bson:"_id" json:"id"`
	ImgUri		string `json:"img_uri"`
	BiodataId	string `json:"biodata_id"`
}

type EducationDegree struct {
	Id 			primitive.ObjectID `bson:"_id" json:"id"`
	Degree		string `json:"degree"`
}

type Education struct {
	Id 			primitive.ObjectID `bson:"_id" json:"id"`
	StartDate	string `json:"start_date"`
	EndDate		string `json:"end_date"`
	SchoolName	string `json:"school_name"`
	Gpa			int32  `json:"gpa"`
	Degree		string `json:"degree"`
	Major		string `json:"major"`
	BiodataId	string `json:"biodata_id"`
}

type BiodataFull struct {
	Biodata 	Biodata
	AboutMe		AboutMe
	Contact 	Contact
	SocialMedia []SocialMedia
	ProfilePicture ProfilePicture
}

type ResponseBiodata struct {
	Status 		int `json:"status"`
	Message 	string `json:"message"`
	Data 		[]BiodataFull
}
