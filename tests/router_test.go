package tests

import (
	"fmt"
	"github.com/michaelwp/resume-backend/servers"
	"strings"
	"testing"
)


var _, _, resp = servers.Server("localhost:8080")

func TestRouterCon(t *testing.T)  {
	fmt.Println("test router connection : ")

	if strings.Compare(resp, "Server running on localhost:8080") != 0 {
		t.Errorf("Not connected to main router")
	}
}
