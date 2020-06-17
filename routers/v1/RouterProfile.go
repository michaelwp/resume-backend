package v1

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/controllers/v1"
)

func RouterProfile(r *mux.Router) {
	r.HandleFunc("/profiles", v1.GetProfile).Methods("GET")
	r.HandleFunc("/profiles", v1.PostBiodata).Methods("POST")

	profiles := r.PathPrefix("/profiles").Subrouter()

	profiles.HandleFunc("/abouts", v1.PostAboutMe).Methods("POST")
	profiles.HandleFunc("/contacts", v1.PostContact).Methods("POST")
	profiles.HandleFunc("/socials", v1.PostSocialMedia).Methods("POST")
	profiles.HandleFunc("/pictures", v1.PostProfilePicture).Methods("POST")
}