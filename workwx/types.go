package workwx

import (
	"github.com/anerg2046/goauth/authtype"
	"github.com/muesli/cache2go"
)

type WorkWx struct {
	conf  *authtype.AuthConf
	cache *cache2go.CacheTable
}

type err struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type RspAccessToken struct {
	err
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIn   int    `json:"expires_in,omitempty"`
}

type RspUserInfo struct {
	err
	UserId string `json:"UserId,omitempty"`
	OpenId string `json:"OpenId,omitempty"`
}

type RspEmployee struct {
	err
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Gender string `json:"gender,omitempty"`
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`
}

type RspOpenId struct {
	err
	OpenId string `json:"openid,omitempty"`
}
