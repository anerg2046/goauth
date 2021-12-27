package goauth

import (
	"strings"

	"github.com/anerg2046/goauth/authtype"
	"github.com/anerg2046/goauth/i"
	"github.com/anerg2046/goauth/workwx"

	"github.com/muesli/cache2go"
)

var cache *cache2go.CacheTable

func init() {
	cache = cache2go.Cache("goauth")
}

func NewGoAuth(platform string, config *authtype.AuthConf) i.GoAuth {
	platform = strings.ToLower(platform)
	switch platform {
	case "workwx":
		return workwx.NewWorkWx(config, cache)
	}
	return nil
}
