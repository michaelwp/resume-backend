package main

import (
	"fmt"
	"github.com/michaelwp/resume-backend/servers"
	"log"
)

func main()  {
	srv, _, resp, _, _ := servers.Server()

	fmt.Println(resp)
	log.Fatal(srv.ListenAndServe())
}
