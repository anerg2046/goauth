package workwx

import (
	"encoding/json"

	e "github.com/anerg2046/goauth/error"
	"github.com/anerg2046/goauth/r"
)

const ApiUri = "https://qyapi.weixin.qq.com/cgi-bin/"

func (auth *WorkWx) getToken() (accessToken RspAccessToken) {
	resp, err := r.HttpClient.R().SetQueryParam("corpid", auth.conf.AppID).SetQueryParam("corpsecret", auth.conf.AppSecret).Get(ApiUri + "gettoken")
	if err != nil {
		panic(&e.GoAuthError{Err: err.Error(), Info: "请求企业微信接口出错-获取Token"})
	}
	json.Unmarshal(resp.Body(), &accessToken)
	return
}

func (auth *WorkWx) getUserInfo(code string) (userInfo RspUserInfo) {
	resp, err := r.HttpClient.R().SetQueryParam("access_token", auth.AccessToken()).SetQueryParam("code", code).Get(ApiUri + "user/getuserinfo")
	if err != nil {
		panic(&e.GoAuthError{Err: err.Error(), Info: "请求企业微信接口出错-获取访问用户身份"})
	}
	json.Unmarshal(resp.Body(), &userInfo)
	return
}

func (auth *WorkWx) getEmployee(userId string) (employee RspEmployee) {
	resp, err := r.HttpClient.R().SetQueryParam("access_token", auth.AccessToken()).SetQueryParam("userid", userId).Get(ApiUri + "user/get")
	if err != nil {
		panic(&e.GoAuthError{Err: err.Error(), Info: "请求企业微信接口出错-读取成员"})
	}
	json.Unmarshal(resp.Body(), &employee)
	return
}
