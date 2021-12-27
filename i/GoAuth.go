package i

import "github.com/anerg2046/goauth/authtype"

type GoAuth interface {
	Platform() string
	AccessToken() string
	GetRedirectUrl() string
	GetUserInfo(code string) authtype.UserInfo
}
