
package user

import (
	"testing"
	"s3-web-browser/server/go/domain/db"
)


func TestTransaction(t *testing.T) {
	conn, err := db.ConnectionForTest()
	if err != nil {
		t.Fatalf("failed test %#v", err)
		return
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		t.Fatalf("failed test %#v", err)
		return
	}

	user := User{Loginid: "loginid", Username: "test", Password: "pass"}

	id, err := Insert(tx, &user)
	if err != nil {
		t.Fatalf("failed test %#v", err)
		tx.Rollback()
		return
	}
	user.ID = int(id)

	userInserted, err := SelectByID(tx, id)
	if user.Loginid != userInserted.Loginid ||
		user.Username != userInserted.Username ||
		"********" != userInserted.Password {
		t.Errorf("Actual: %s, but excepted: %s", user.String(), userInserted.String())
	}

	user.Loginid = "hoge"
	user.Username = "huga"
	_, err = UpdateByID(tx, &user)
	if err != nil {
		t.Fatalf("failed test %#v", err)
		tx.Rollback()
		return
	}
	_, err = UpdatePasswordByID(tx, &user)
	if err != nil {
		t.Fatalf("failed test %#v", err)
		tx.Rollback()
		return
	}

	userUpdateed, err := SelectByID(tx, id)
	if err != nil {
		t.Fatalf("failed test %#v", err)
		tx.Rollback()
		return
	}
	if user.Loginid != userUpdateed.Loginid ||
		user.Username != userUpdateed.Username ||
		"********" != userInserted.Password {
		t.Errorf("Actual: %s, but excepted: %s", user.String(), userUpdateed.String())
	}
	
	_, err = SelectAll(tx)
	if err != nil {
		t.Fatalf("failed test %#v", err)
		tx.Rollback()
		return
	}

	tx.Rollback()
}
