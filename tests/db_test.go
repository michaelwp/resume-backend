package tests

import (
	"fmt"
	"github.com/michaelwp/resume-backend/db"
	"strings"
	"testing"
)

func TestDbCon(t *testing.T){
	fmt.Println("test db connection : ")

	_, status := db.DbCon("db_resume")

	if strings.Compare(status, "Connected to MongoDB!") != 0 {
		t.Errorf("Not connected to mongodb")
	}
}
