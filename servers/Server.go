package servers

import (
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/routers"
	"net/http"
	"strings"
	"time"
)

func Server() (*http.Server, *mux.Router, string, string, string)  {
	port := ":8080"
	host := "localhost"
	resp := strings.Join([]string{"Server", host, "running and listening on port", port}, " ")
	router := mux.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr: strings.Join([]string{host, port},""),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	//connecting to router
	routers.MainRouter(router)

	return srv, router, resp, host, port
}
