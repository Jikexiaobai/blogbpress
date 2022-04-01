package sms

// Engine 短信引擎
type Engine interface {
	Send() error
}
