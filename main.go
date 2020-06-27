package main

import (
	"fmt"
	"github.com/michaelwp/resume-backend/helpers"
	"github.com/michaelwp/resume-backend/servers"
	"log"
	"strings"
)

func main()  {
	port := helpers.LoadEnv("HOST_PORT")
	host := helpers.LoadEnv("HOST_DEV")

	h := strings.Join([]string{host, port},"")

	srv, _, resp := servers.Server(h)

	fmt.Println(resp)
	log.Fatal(srv.ListenAndServe())
}
