package routers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/routers/v1"
)

func RouterCon() (*mux.Router, string, string)  {
	port := ":8080"
	resp := "Server running and listening on port"
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	//api version 1
	ver1 := api.PathPrefix("/v1").Subrouter()
	v1.RouterProfile(ver1)

	return router, resp, port
}
