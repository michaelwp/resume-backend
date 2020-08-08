package v1

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/db"
	"github.com/michaelwp/resume-backend/helpers"
	"github.com/michaelwp/resume-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
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
	Profile [GET]
	http://localhost:8080/api/v1/profiles/{id}
	=====================================================
*/
func GetProfileById(w http.ResponseWriter, r *http.Request)  {
	var f []models.BiodataFull
	vars := mux.Vars(r)
	f = append(f, findBiodata(vars["id"]))
	responseProfile(w, 1, "Profile Data", f, http.StatusCreated)
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