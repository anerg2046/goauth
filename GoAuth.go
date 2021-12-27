package goauth

import (
	"goauth/goauthconf"
	"goauth/i"
	"goauth/workwx"
	"strings"

	"github.com/muesli/cache2go"
)

var cache *cache2go.CacheTable

func init() {
	cache = cache2go.Cache("goauth")
}

func NewGoAuth(platform string, config *goauthconf.AuthConf) i.GoAuth {
	platform = strings.ToLower(platform)
	switch platform {
	case "workwx":
		return workwx.NewWorkWx(config, cache)
	}
	return nil
}
