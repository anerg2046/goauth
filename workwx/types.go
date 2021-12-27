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

type RspAccessToken struct {
	Errcode     int           `json:"errcode,omitempty"`
	Errmsg      string        `json:"errmsg,omitempty"`
	AccessToken string        `json:"access_token,omitempty"`
	ExpiresIn   time.Duration `json:"expires_in,omitempty"`
}
