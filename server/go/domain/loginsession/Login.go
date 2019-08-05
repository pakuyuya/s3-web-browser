package login

import "errors"

// Logininfo is a struct for manage login session.
type Logininfo struct {
	UserID    string
	UserName  string
	IsEnabled bool
}

// Auth is a function that authentication user.
func Auth(loginid string, password string) (*Logininfo, error) {
	// TODO:

	if loginid != "user" || password != "pass" {
		return nil, errors.New("ログインIDまたはパスワードが違います")
	}

	info := Logininfo{
		UserID:    loginid,
		UserName:  "テストユーザー",
		IsEnabled: true,
	}

	return &info, nil
}

// SessionKey is the key of login infomation for sessions.
const SessionKey = "Login"
