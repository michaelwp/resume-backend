package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/michaelwp/resume-backend/db"
	"github.com/michaelwp/resume-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

/*
	=====================================================
	Home [GET]
	http://localhost:8080/profile
	=====================================================
*/
func GetProfile(w http.ResponseWriter, r *http.Request)  {
	var responseBiodata models.ResponseBiodata
	var biodata []models.Biodata
	var biodataFull models.BiodataFull
	var biodataFullRes []models.BiodataFull

	cur, err := db.DbCon().Collection("tbl_biodata").Find(context.TODO(), bson.D{})
	if err != nil {log.Fatal(err)}

	for cur.Next(context.TODO()) {
		var el = models.Biodata{}
		err = cur.Decode(&el)
		if err != nil {log.Fatal(err)}
		biodata = append(biodata, el)
	}

	if err = cur.Err(); err!= nil {log.Fatal(err)}

	err = cur.Close(context.TODO())
	if err != nil {log.Fatal(err)}

	//biodataFull
	for _, b := range biodata {
		biodataFull.Biodata = b
		biodataFullRes = append(biodataFullRes, biodataFull)
	}

	//Response
	responseBiodata.Status = 1
	responseBiodata.Message = "Biodata"
	responseBiodata.Data = biodataFullRes

	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(responseBiodata)
	if err != nil {log.Fatal(err)}
}

/*
	=====================================================
	Home [POST]
	http://localhost:8080/profile
	=====================================================
*/
func PostBiodata(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintln(w, "Post Profile")
}
