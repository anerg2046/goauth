package workwx

import (
	"encoding/json"

	e "github.com/anerg2046/goauth/error"
	"github.com/anerg2046/goauth/r"
)

const ApiUri = "https://qyapi.weixin.qq.com/cgi-bin"

func (auth *WorkWx) getToken() RspAccessToken {
	resp, err := r.HttpClient.R().SetQueryParam("corpid", auth.conf.AppID).SetQueryParam("corpsecret", auth.conf.AppSecret).Get(ApiUri + "/gettoken")
	if err != nil {
		panic(&e.GoAuthError{Err: err.Error(), Info: "请求企业微信接口出错"})
	}
	var accessToken RspAccessToken
	json.Unmarshal(resp.Body(), &accessToken)
	return accessToken
}
