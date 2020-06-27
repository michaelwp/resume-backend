package v1

import (
	"context"
	"encoding/json"
	"github.com/michaelwp/resume-backend/db"
	"github.com/michaelwp/resume-backend/helpers"
	"github.com/michaelwp/resume-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"strings"
)

var dbCon, _ = db.DbCon("db_resume")

/*
	=====================================================
	Profile [GET]
	http://localhost:8080/api/v1/profiles
	=====================================================
*/
func GetProfile(w http.ResponseWriter, r *http.Request)  {
	var biodata []models.Biodata
	var biodataFull models.BiodataFull
	var biodataFullRes []models.BiodataFull

	cur, err := dbCon.Collection("tbl_biodata").Find(context.TODO(), bson.D{})
	if err != nil {log.Fatal(err)}

	for cur.Next(context.TODO()) {
		var el = models.Biodata{}
		err = cur.Decode(&el)
		if err != nil {log.Fatal(err)}
		el.BirthDate = helpers.DateFormat(el.BirthDate, "US")
		biodata = append(biodata, el)
	}

	if err = cur.Err(); err!= nil {log.Fatal(err)}

	err = cur.Close(context.TODO())
	if err != nil {log.Fatal(err)}

	//biodataFull
	for _, b := range biodata {
		biodataFull.AboutMe = findAboutMe(b.Id.Hex())
		biodataFull.Contact = findContact(b.Id.Hex())
		biodataFull.SocialMedia = findSocialMedia(b.Id.Hex())
		biodataFull.ProfilePicture = findProfilePicture(b.Id.Hex())
		biodataFull.Biodata = b
		biodataFullRes = append(biodataFullRes, biodataFull)
	}

	responseProfile(w, 1, "Profile Data", biodataFullRes, http.StatusOK)
}

/*
	=====================================================
	Find Biodata by Object Id
	=====================================================
*/
func findBiodata(id string) models.BiodataFull {
	var b models.Biodata
	var f models.BiodataFull

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {log.Fatal(err)}

	filter := bson.M{"_id": oid}
	err = dbCon.Collection("tbl_biodata").FindOne(context.TODO(), filter).Decode(&b)
	if err != nil {log.Println(err)}

	f.AboutMe = findAboutMe(id)
	f.Contact = findContact(id)
	f.SocialMedia = findSocialMedia(id)
	f.ProfilePicture = findProfilePicture(id)
	f.Biodata = b

	return f
}

/*
	=====================================================
	Find About Me by biodata Id
	=====================================================
*/
func findAboutMe(id string) models.AboutMe{
	var a models.AboutMe
	filter := bson.M{"biodataid": id}

	err := dbCon.Collection("tbl_about_me").FindOne(context.TODO(), filter).Decode(&a)
	if err != nil {log.Println(err)}

	return a
}

/*
	=====================================================
	Find Contact by profile Id
	=====================================================
*/
func findContact(id string) models.Contact{
	var c models.Contact
	filter := bson.M{"biodataid": id}

	err := dbCon.Collection("tbl_contact").FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {log.Println(err)}

	return c
}

/*
	=====================================================
	Find Social Media by profile Id
	=====================================================
*/
func findSocialMedia(id string) []models.SocialMedia{
	var socialMedia []models.SocialMedia
	filter := bson.M{"biodataid": id}

	cur, err := dbCon.Collection("tbl_social_media").Find(context.TODO(), filter)
	if err != nil {log.Println(err)}

	for cur.Next(context.TODO()) {
		var el = models.SocialMedia{}
		err = cur.Decode(&el)
		if err != nil {log.Fatal(err)}
		socialMedia = append(socialMedia, el)
	}

	if err = cur.Err(); err!= nil {log.Fatal(err)}

	err = cur.Close(context.TODO())
	if err != nil {log.Fatal(err)}

	return socialMedia
}

/*
	=====================================================
	Find Profile Picture by profile Id
	=====================================================
*/
func findProfilePicture(id string) models.ProfilePicture {
	var pp models.ProfilePicture
	filter := bson.M{"biodataid": id}

	err := dbCon.Collection("tbl_profile_picture").FindOne(context.TODO(), filter).Decode(&pp)
	if err != nil {log.Println(err)}

	return pp
}

/*
	=====================================================
	Home [POST]
	http://localhost:8080/api/v1/profiles
	request Body = {
		"first_name" : "John"
		"last_name" : "Doe"
		...
		"zip_code" : "123456"
	}
	=====================================================
*/
func PostProfile(w http.ResponseWriter, r *http.Request)  {
	var b models.Biodata
	var f []models.BiodataFull

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	dataInput := bson.M{
		"firstname": strings.ToLower(b.FirstName),
		"lastname": strings.ToLower(b.LastName),
		"birthplace": strings.ToLower(b.BirthPlace),
		"birthdate": strings.ToLower(b.BirthDate),
		"address": strings.ToLower(b.Address),
		"city": strings.ToLower(b.City),
		"province": strings.ToLower(b.Province),
		"state": strings.ToLower(b.State),
		"zipcode": strings.ToLower(b.ZipCode),
	}

	res, err := dbCon.Collection("tbl_biodata").InsertOne(context.TODO(), dataInput)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)
	f = append(f, findBiodata(oid.Hex()))

	responseProfile(w, 1, "Data saved successfully", f, http.StatusCreated)
}

/*
	=====================================================
	Home [POST]
	http://localhost:8080/api/v1/profiles/abouts
	request Body = {
		"about_me": "Software Engineer",
		"biodata_id": "123456"
	}
	=====================================================
*/
func PostAbout(w http.ResponseWriter, r *http.Request)  {
	var a models.AboutMe
	var f []models.BiodataFull

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	dataInput := bson.M{
		"aboutme": strings.ToLower(a.AboutMe),
		"biodataid": a.BiodataId,
	}

	_, err = dbCon.Collection("tbl_about_me").InsertOne(context.TODO(), dataInput)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	f = append(f, findBiodata(a.BiodataId))

	responseProfile(w, 1, "Data saved successfully", f, http.StatusCreated)
}

/*
	=====================================================
	Home [POST]
	http://localhost:8080/api/v1/profiles/contacts
	request Body = {
		"phone_number": "+62 123",
        "email": "mail@mail.com",
        "biodata_id": "123456"
	}
	=====================================================
*/
func PostContact(w http.ResponseWriter, r *http.Request)  {
	var c models.Contact
	var f []models.BiodataFull

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	dataInput := bson.M{
		"phonenumber": strings.ToLower(c.PhoneNumber),
		"email": c.Email,
		"biodataid": c.BiodataId,
	}

	_, err = dbCon.Collection("tbl_contact").InsertOne(context.TODO(), dataInput)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	f = append(f, findBiodata(c.BiodataId))

	responseProfile(w, 1, "Data saved successfully", f, http.StatusCreated)
}

/*
	=====================================================
	Home [POST]
	http://localhost:8080/api/v1/profiles/socials
	request Body = {
		"name": "linkedin",
        "link": "http://www.linkedin.com/in/michael-wenceslaus",
        "biodata_id": "123456"
	}
	=====================================================
*/
func PostSocial(w http.ResponseWriter, r *http.Request)  {
	var s models.SocialMedia
	var f []models.BiodataFull

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	dataInput := bson.M{
		"name":strings.ToLower(s.Name),
		"link":strings.ToLower(s.Link),
		"biodataid":s.BiodataId,
	}

	_, err = dbCon.Collection("tbl_social_media").InsertOne(context.TODO(), dataInput)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	f = append(f, findBiodata(s.BiodataId))

	responseProfile(w, 1, "Data saved successfully", f, http.StatusCreated)
}

/*
	=====================================================
	Home [POST]
	http://localhost:8080/api/v1/profiles/pictures
	request Body = {
		"img_uri": "https://i.imgur.com/pgPfRVW.jpg",
        "biodata_id": "123456"
	}
	=====================================================
*/
func PostPicture(w http.ResponseWriter, r *http.Request)  {
	var p models.ProfilePicture
	var f []models.BiodataFull

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	dataInput := bson.M{
		"imguri":strings.ToLower(p.ImgUri),
		"biodataid":p.BiodataId,
	}

	_, err = dbCon.Collection("tbl_profile_picture").InsertOne(context.TODO(), dataInput)
	if err != nil {
		responseProfile(w, 0, err.Error(), f, http.StatusBadRequest)
		return
	}

	f = append(f, findBiodata(p.BiodataId))

	responseProfile(w, 1, "Data saved successfully", f, http.StatusCreated)
}


