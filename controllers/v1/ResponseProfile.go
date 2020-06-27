package v1

import (
	"encoding/json"
	"github.com/michaelwp/resume-backend/models"
	"log"
	"net/http"
)

func responseProfile(w http.ResponseWriter, s int, m string, b []models.BiodataFull, h int) {
	var r models.ResponseBiodata

	r.Status = s
	r.Message = m
	r.Data = b

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(h)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {log.Fatal(err)}
}
