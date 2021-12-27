package authtype

type UserInfo struct {
	OpenId  string `json:"open_id,omitempty"`
	UnionId string `json:"union_id,omitempty"`
	Source  string `json:"source,omitempty"`
	Nick    string `json:"nick,omitempty"`
	Gender  string `json:"gender,omitempty"`
	Avatar  string `json:"avatar,omitempty"`
	Email   string `json:"email,omitempty"`
}
