package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

/*
	=====================================================
	Profile [DELETE]
	http://localhost:8080/api/v1/profiles/{id}
	=====================================================
*/
func DeleteProfileById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	result := deleteBiodata(vars["id"])
	res := make(map[string]string)

	res["Status"] = "1"
	res["Message"] = result

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {log.Fatal(err)}
}

/*
	=====================================================
	Delete Biodata by Object Id
	=====================================================
*/
func deleteBiodata(id string) string {
	var result string

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {log.Fatal(err)}

	filter := bson.M{"_id": oid}
	res, err := dbCon.Collection("tbl_biodata").DeleteOne(context.TODO(), filter)
	if err != nil {log.Println(err)}

	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		// set the result
		result = "Profile not found"
	} else {
		// set the result
		result = "Profile Successfully delete"
	}

	// delete others
	deleteAboutMe(id)
	deleteContact(id)
	deleteSocialMedia(id)
	deleteProfilePicture(id)

	return result
}

/*
	=====================================================
	Delete About Me by Object Id
	=====================================================
*/
func deleteAboutMe(id string) {
	var result string
	filter := bson.M{"biodataid": id}

	res, err := dbCon.Collection("tbl_about_me").DeleteOne(context.TODO(), filter)
	if err != nil {log.Println(err)}

	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		// set the result
		result = "About Me not found"
	} else {
		// set the result
		result = "About Me Successfully delete"
	}

	fmt.Println(result)
}

/*
	=====================================================
	Delete Contact by profile Id
	=====================================================
*/
func deleteContact(id string){
	var result string
	filter := bson.M{"biodataid": id}

	res, err := dbCon.Collection("tbl_contact").DeleteOne(context.TODO(), filter)
	if err != nil {log.Println(err)}

	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		// set the result
		result = "Contact not found"
	} else {
		// set the result
		result = "Contact Successfully delete"
	}

	fmt.Println(result)
}

/*
	=====================================================
	Delete Social Media by profile Id
	=====================================================
*/
func deleteSocialMedia(id string) {
	var result string
	filter := bson.M{"biodataid": id}

	res, err := dbCon.Collection("tbl_social_media").DeleteOne(context.TODO(), filter)
	if err != nil {log.Println(err)}

	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		// set the result
		result = "Social media not found"
	} else {
		// set the result
		result = "Social media Successfully delete"
	}

	fmt.Println(result)
}

/*
	=====================================================
	Delete Profile Picture by profile Id
	=====================================================
*/
func deleteProfilePicture(id string) {
	var result string
	filter := bson.M{"biodataid": id}

	res, err := dbCon.Collection("tbl_profile_picture").DeleteOne(context.TODO(), filter)
	if err != nil {log.Println(err)}

	// Check if the response is 'nil'
	if res.DeletedCount == 0 {
		// set the result
		result = "Profile picture not found"
	} else {
		// set the result
		result = "Profile picture Successfully delete"
	}

	fmt.Println(result)
}