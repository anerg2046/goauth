package workwx

import (
	"time"

	"github.com/anerg2046/goauth/goauthconf"

	"github.com/muesli/cache2go"
)

type WorkWx struct {
	conf  *goauthconf.AuthConf
	cache *cache2go.CacheTable
}

type err struct {
	Errcode int    `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

type RspAccessToken struct {
	err
	AccessToken string        `json:"access_token,omitempty"`
	ExpiresIn   time.Duration `json:"expires_in,omitempty"`
}

type RspUserInfo struct {
	err
	UserId string `json:"user_id,omitempty"`
	OpenId string `json:"open_id,omitempty"`
}
