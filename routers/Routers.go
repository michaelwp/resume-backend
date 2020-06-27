package routers

import (
	"github.com/gorilla/mux"
	v1 "github.com/michaelwp/resume-backend/routers/v1"
)

func MainRouter(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	routerVersion1(api)
}

/* api version 1 */
func routerVersion1(r *mux.Router) {
	ver1 := r.PathPrefix("/v1").Subrouter()
	v1.RouterProfile(ver1)
}
