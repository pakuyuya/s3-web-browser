package login

import (
	"database/sql"
	"fmt"
	"s3-web-browser/server/go/domain/db"
	"s3-web-browser/server/go/domain/user"

	"errors"
	_ "github.com/lib/pq"
)

// Logininfo is a struct for manage login session.
type Logininfo struct {
	UserID    string
	UserName  string
	IsEnabled bool
}

// Auth is a function that authentication user.
func Auth(loginid string, password string) (*Logininfo, error) {

	conn, err := db.Connection()
	if err != nil {
		return nil, err
	}
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	user, err := user.SelectForAuth(loginid, password)
	if user == nil {
		return nil, errors.New("ログインIDまたはパスワードが違います")
	}

	info := Logininfo{
		UserID:    user.Loginid,
		UserName:  user.Username,
		IsEnabled: true,
	}

	return &info, nil
}

// SessionKey is the key of login infomation for sessions.
const SessionKey = "Login"
