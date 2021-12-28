package authtype

type SnsSource uint8

const (
	_ SnsSource = iota
	WORKWX
	WECHAT
)

func (s SnsSource) String() string {
	switch s {
	case WORKWX:
		return "企业微信"
	case WECHAT:
		return "微信"
	default:
		return "未知"
	}
}
