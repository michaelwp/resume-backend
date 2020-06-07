package routers

import "github.com/gorilla/mux"

func RouterCon() *mux.Router  {
	router := mux.NewRouter()

	routerProfile(router)

	return router
}
