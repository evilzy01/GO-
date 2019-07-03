package dbops

import (
	"fmt"
	"testing"
)

//////////////////////////////////////////////

func Testmain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestUserWorkflow(t *testing.T) {
	t.Run("add", testAddUsers)
	t.Run("get", testGetUsers)
	t.Run("del", testDelUsers)
	t.Run("Reget", testRegetUsers)
}

////////////////////////////////////////
func testAddUsers(t *testing.T) {
	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of add users: %v", err)
	}
}

func testGetUsers(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of get users")
	}
}

func testDelUsers(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of del users: %v", err)
	}
}

func testRegetUsers(t *testing.T) {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of Reget users: %v", err)
	}
	if pwd != "" {
		t.Errorf("Del Failed")
	}
}
