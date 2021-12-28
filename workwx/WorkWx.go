package workwx

import (
	"net/url"

	"github.com/anerg2046/go-pkg/logrus"
	"github.com/anerg2046/goauth/authtype"
	e "github.com/anerg2046/goauth/error"
	"github.com/anerg2046/goauth/i"

	"github.com/muesli/cache2go"
)

func NewWorkWx(config *authtype.AuthConf, cache *cache2go.CacheTable) i.GoAuth {
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
		logrus.Log.Info("请求获取token")
		token := auth.getToken()
		if token.Errcode == 0 {
			accessToken = token.AccessToken
			auth.cache.Add(cacheKey, token.ExpiresIn, &accessToken)
		} else {
			panic(&e.GoAuthError{Err: token.Errmsg, Info: "获取企业微信Token失败"})
		}
	} else {
		logrus.Log.Info("缓存获取token")
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

func (auth *WorkWx) GetUserInfo(code string) (userInfo authtype.UserInfo) {
	uinfo := auth.getUserInfo(code)
	if uinfo.UserId != "" {
		employee := auth.getEmployee(uinfo.UserId)
		userInfo.OpenId = auth.getOpenId(uinfo.UserId).OpenId
		userInfo.Avatar = employee.Avatar
		userInfo.Email = employee.Email
		userInfo.Mobile = employee.Mobile
		userInfo.Gender = parseGender(employee.Gender)
		userInfo.Nick = employee.Name
		userInfo.IsEmployee = true
	} else {
		userInfo.OpenId = uinfo.OpenId
		userInfo.IsEmployee = false
	}
	userInfo.Source = "workwx"
	return
}

func parseGender(gender string) string {
	switch gender {
	case "1":
		return "m"
	case "2":
		return "f"
	}
	return "n"
}
