package workwx

import (
	e "goauth/error"
	"goauth/goauthconf"
	"goauth/i"

	"github.com/muesli/cache2go"
)

func NewWorkWx(config *goauthconf.AuthConf, cache *cache2go.CacheTable) i.GoAuth {
	return &WorkWx{
		conf:  config,
		cache: cache,
	}
}

func (auth *WorkWx) Platform() string {
	return "workwx"
}

func (auth *WorkWx) AccessToken() string {
	var accessToken string
	cacheKey := auth.Platform() + "accessToken"
	res, err := auth.cache.Value(cacheKey)
	if err != nil {
		token := auth.getToken()
		if token.Errcode == 0 {
			accessToken = token.AccessToken
			auth.cache.Add(cacheKey, token.ExpiresIn, &accessToken)
		} else {
			panic(&e.GoAuthError{Err: token.Errmsg, Info: "获取企业微信Token失败"})
		}
	} else {
		accessToken = *res.Data().(*string)
	}
	return accessToken
}
