package v1

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/controllers/v1"
)

func RouterProfile(r *mux.Router) {
	// GET PROFILES
	r.HandleFunc("/profiles", v1.GetProfile).Methods("GET")
	r.HandleFunc("/profiles/{id}", v1.GetProfileById).Methods("GET")

	profiles := r.PathPrefix("/profiles").Subrouter()

	// POST PROFILES
	r.HandleFunc("/profiles", v1.PostProfile).Methods("POST")
	profiles.HandleFunc("/abouts", v1.PostAbout).Methods("POST")
	profiles.HandleFunc("/contacts", v1.PostContact).Methods("POST")
	profiles.HandleFunc("/socials", v1.PostSocial).Methods("POST")
	profiles.HandleFunc("/pictures", v1.PostPicture).Methods("POST")

	// DELETE PROFILES
	r.HandleFunc("/profiles/{id}", v1.DeleteProfileById).Methods("DELETE")
}
