package servers

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/michaelwp/resume-backend/routers"
	"net/http"
	"os"
	"strings"
	"time"
)

func Server() (*http.Server, *mux.Router, string, string, string)  {
	port := ":8080"
	host := "localhost"
	resp := strings.Join([]string{"Server", host, "running and listening on port", port}, " ")
	router := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	srv := &http.Server{
		Handler: loggedRouter,
		Addr: strings.Join([]string{host, port},""),
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	//connecting to router
	routers.MainRouter(router)

	return srv, router, resp, host, port
}
