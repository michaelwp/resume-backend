package servers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/middlewares"
	"github.com/michaelwp/resume-backend/routers"
	"net/http"
	"strings"
	"time"
)

func Server(h string) (*http.Server, *mux.Router, string)  {
	resp := strings.Join([]string{"Server running on", h}, " ")
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	loggedRouter := middlewares.RouterLogger(router)
	srv := &http.Server{
		Handler: loggedRouter,
		Addr: h,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	//connecting to router
	routers.MainRouter(router)

	return srv, router, resp
}
