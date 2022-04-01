package notice

type WSData struct {
	Code    string      `json:"code,omitempty"`
	Token   string      `json:"token,omitempty"`
	Message interface{} `json:"message,omitempty"`
	WsKey   string      `json:"-"`
}
