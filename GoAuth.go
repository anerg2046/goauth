package goauth

import (
	"github.com/anerg2046/goauth/authtype"
	"github.com/anerg2046/goauth/i"
	"github.com/anerg2046/goauth/workwx"

	"github.com/muesli/cache2go"
)

var cache *cache2go.CacheTable

func init() {
	cache = cache2go.Cache("goauth")
}

func NewGoAuth(source authtype.SnsSource, config *authtype.AuthConf) i.GoAuth {
	switch source {
	case authtype.WORKWX:
		return workwx.NewWorkWx(config, cache)
	}
	return nil
}
