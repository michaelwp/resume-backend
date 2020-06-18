package tests

import (
	"fmt"
	"github.com/michaelwp/resume-backend/routers"
	"testing"
)

func TestRouterCon(t *testing.T)  {
	fmt.Println("test router connection : ")
	_, resp, port := routers.RouterCon()

	serverResp := resp + " " + port

	if serverResp != "Server running and listening on port :8080" {
		t.Errorf("Server error")
	}
}
