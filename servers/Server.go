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

func Server(h string) (*http.Server, *mux.Router, string)  {
	resp := strings.Join([]string{"Server running on", h}, " ")
	router := mux.NewRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
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
