package loginsession

import (
	"database/sql"
	"errors"
	"s3-web-browser/server/go/domain/user"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// Logininfo is a struct for manage login session.
type Logininfo struct {
	UserID      string                 `json:"userid"`
	UserName    string                 `json:"username"`
	IsEnabled   bool                   `json:"isEnabled"`
	Permissions map[string]interface{} `json:"permissions"`
}

// Auth is a function that authentication user.
func Auth(tx *sql.Tx, loginid string, password string) (*Logininfo, error) {
	user, _ := user.SelectForAuth(tx, loginid, password)
	if user == nil {
		return nil, errors.New("ログインIDまたはパスワードが違います")
	}

	info := Logininfo{
		UserID:      user.Loginid,
		UserName:    user.Username,
		IsEnabled:   true,
		Permissions: user.Permissions,
	}

	return &info, nil
}

// SessionKey is the key of login infomation for sessions.
const SessionKey = "Login"
