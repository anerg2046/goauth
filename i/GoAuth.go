package i

type GoAuth interface {
	Platform() string
	AccessToken() string
}
