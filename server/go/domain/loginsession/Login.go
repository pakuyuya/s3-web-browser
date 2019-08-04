package login

import "errors"

type LoginInfo struct {
	UserId    string
	UserName  string
	IsEnabled bool
}

func Auth(loginid string, password string) (*LoginInfo, error) {
	// TODO:

	if loginid != "user" || password != "pass" {
		return nil, errors.New("ログインIDまたはパスワードが違います")
	}

	info := LoginInfo{
		UserId:    loginid,
		UserName:  "テストユーザー",
		IsEnabled: true,
	}

	return &info, nil
}

const SessionKey = "Login"
