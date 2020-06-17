package main

import (
	"fmt"
	"github.com/michaelwp/resume-backend/routers"
	"log"
	"net/http"
)

func main()  {
	router, resp, port := routers.RouterCon()

	fmt.Println(resp, port)
	log.Fatal(http.ListenAndServe(port, router))
}
