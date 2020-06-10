package routers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/controllers"
)

func routerProfile(r *mux.Router) {
	r.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
	r.HandleFunc("/profile", controllers.PostBiodata).Methods("POST")
	r.HandleFunc("/aboutme", controllers.PostAboutMe).Methods("POST")
	r.HandleFunc("/contact", controllers.PostContact).Methods("POST")
	r.HandleFunc("/socialmedia", controllers.PostSocialMedia).Methods("POST")
	r.HandleFunc("/profilepicture",controllers.PostProfilePicture).Methods("POST")
}
