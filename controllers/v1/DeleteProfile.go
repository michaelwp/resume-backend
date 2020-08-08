package v1

import (
	"context"
	"encoding/json"
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
		result = "Document not found"
	} else {
		// set the result
		result = "Successfully delete"
	}

	return result
}