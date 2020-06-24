package tests

import (
	"fmt"
	"github.com/michaelwp/resume-backend/routers"
	v1 "github.com/michaelwp/resume-backend/routers/v1"
	"strings"
	"testing"
)

var router, resp, port = routers.RouterCon()
var api = router.PathPrefix("/api").Subrouter()
var routerResp string

//api version 1
var ver1 = api.PathPrefix("/v1").Subrouter()

func TestRouterCon(t *testing.T)  {
	fmt.Println("test router connection : ")

	serverResp := resp + " " + port

	if strings.Compare(serverResp, "Server running and listening on port :8080") == 1 {
		t.Errorf("Not connected to main router")
	}
}

func TestRouterProfile(t *testing.T) {
	fmt.Println("test router profile version 1 : ")

	routerResp = v1.RouterProfile(ver1)

	if strings.Compare("router profile version 1", routerResp) == 1 {
		t.Errorf("Not connected to router profile version 1")
	}
}
