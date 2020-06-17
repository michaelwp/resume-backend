package routers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/routers/v1"
)

func RouterCon() *mux.Router  {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()

	//api version 1
	ver1 := api.PathPrefix("/v1").Subrouter()
	v1.RouterProfile(ver1)

	return router
}
