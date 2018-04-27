package table

import (
	"log"
	"testing"

	"gopromo/table/user"
)

func TestUserGetById(t *testing.T) {
	userTab := userTable.New()
	user, err := userTab.GetById(1)
	if err != nil {
		t.Errorf("err: %s", err.Error())
	} else {
		log.Println("user id=1 name=", user.Name)
	}
}
