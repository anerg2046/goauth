package workwx

import (
	"net/url"

	e "github.com/anerg2046/goauth/error"
	"github.com/anerg2046/goauth/goauthconf"
	"github.com/anerg2046/goauth/i"

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

func (auth *WorkWx) GetRedirectUrl() string {
	var uri url.URL
	q := uri.Query()
	if auth.conf.Scan {
		q.Add("appid", auth.conf.AppID)
		q.Add("agentid", auth.conf.AgentID)
		q.Add("redirect_uri", auth.conf.Callback)
		return "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?" + q.Encode()
	} else {
		q.Add("appid", auth.conf.AppID)
		q.Add("scope", "snsapi_base")
		q.Add("response_type", "code")
		q.Add("redirect_uri", auth.conf.Callback)
		return "https://open.weixin.qq.com/connect/oauth2/authorize?" + q.Encode() + "#wechat_redirect"
	}
}
