package workwx

import (
	"encoding/json"

	e "github.com/anerg2046/goauth/error"
	"github.com/anerg2046/goauth/r"
)

const ApiUri = "https://qyapi.weixin.qq.com/cgi-bin/"

func (auth *WorkWx) getToken() (rsp RspAccessToken) {
	detail := "获取Token"
	resp, err := r.HttpClient.R().SetQueryParam("corpid", auth.conf.AppID).SetQueryParam("corpsecret", auth.conf.AppSecret).Get(ApiUri + "gettoken")
	checkErr(err, detail)
	json.Unmarshal(resp.Body(), &rsp)
	checkRspErr(rsp.err, detail)
	return
}

func (auth *WorkWx) getUserInfo(code string) (rsp RspUserInfo) {
	detail := "获取访问用户身份"
	resp, err := r.HttpClient.R().SetQueryParam("access_token", auth.AccessToken()).SetQueryParam("code", code).Get(ApiUri + "user/getuserinfo")
	checkErr(err, detail)
	json.Unmarshal(resp.Body(), &rsp)
	checkRspErr(rsp.err, detail)
	return
}

func (auth *WorkWx) getEmployee(userId string) (rsp RspEmployee) {
	detail := "读取成员"
	resp, err := r.HttpClient.R().SetQueryParam("access_token", auth.AccessToken()).SetQueryParam("userid", userId).Get(ApiUri + "user/get")
	checkErr(err, detail)
	json.Unmarshal(resp.Body(), &rsp)
	checkRspErr(rsp.err, detail)
	return
}

func (auth *WorkWx) getOpenId(userId string) (rsp RspOpenId) {
	detail := "userid转换openid"
	resp, err := r.HttpClient.R().SetQueryParam("access_token", auth.AccessToken()).SetBody(map[string]string{"userid": userId}).Post(ApiUri + "user/convert_to_openid")
	checkErr(err, detail)
	json.Unmarshal(resp.Body(), &rsp)
	checkRspErr(rsp.err, detail)
	return
}

func checkErr(err error, detail string) {
	if err != nil {
		panic(&e.GoAuthError{Err: err.Error(), Info: "请求企业微信接口出错-" + detail})
	}
}

func checkRspErr(err err, detail string) {
	if err.Errcode != 0 {
		panic(&e.GoAuthError{Err: err.Errmsg, Info: "请求企业微信接口出错-" + detail})
	}
}
