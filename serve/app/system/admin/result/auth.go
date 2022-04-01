package result

// auth
type TokenInfo struct {
	UserId  int64
	Account string
}
type LoginInfo struct {
	NickName     string `json:"nick_name"`
	Avatar       string `json:"avatar"`
	Token        string `json:"token"`
	VerifyStatus int    `json:"verify_status"`
}
