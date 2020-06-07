package routers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/controllers"
)

func routerProfile(r *mux.Router) {
	r.HandleFunc("/profile", controllers.GetProfile).Methods("GET")
	r.HandleFunc("/profile", controllers.PostBiodata).Methods("POST")
}
