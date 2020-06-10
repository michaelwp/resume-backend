package main

import (
	"fmt"
	"github.com/michaelwp/resume-backend/routers"
	"log"
	"net/http"
)

func main()  {
	port := ":8080"
	router := routers.RouterCon()

	fmt.Println("Server running and listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
