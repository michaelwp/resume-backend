package tests

import (
	"fmt"
	"github.com/michaelwp/resume-backend/servers"
	"strings"
	"testing"
)

var srv, router, resp, host, port = servers.Server()
var api = router.PathPrefix("/api").Subrouter()

//api version 1
var ver1 = api.PathPrefix("/v1").Subrouter()

func TestRouterCon(t *testing.T)  {
	fmt.Println("test router connection : ")

	if strings.Compare(resp, "Server localhost running and listening on port :8080") < 0 {
		t.Errorf("Not connected to main router")
	}
}
