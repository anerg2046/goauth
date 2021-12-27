package error

import "github.com/anerg2046/go-pkg/logrus"

type GoAuthError struct {
	Err  string
	Info string
}

func (e *GoAuthError) Error() string {
	logrus.Log.Error(e.Err)
	return e.Info
}
