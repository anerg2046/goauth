package r

import (
	"net/http"
	"time"

	"github.com/anerg2046/go-pkg/logrus"
	"github.com/go-resty/resty/v2"
)

var HttpClient = resty.New()

func init() {
	HttpClient.SetLogger(logrus.Log)
	HttpClient.SetRetryCount(2)
	HttpClient.SetRetryWaitTime(100 * time.Millisecond)
	HttpClient.SetRetryMaxWaitTime(20 * time.Second)
	HttpClient.SetTimeout(5 * time.Second)
	HttpClient.AddRetryCondition(
		func(r *resty.Response, err error) bool {
			if r.StatusCode() != http.StatusOK || r == nil {
				return true
			}
			return false
		},
	)
}
