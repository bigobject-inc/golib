package websocket

// Connection Websocket connection
type Connection interface {
	GetID() string
	IsClosed() bool
	SendMessage(message Message) error
	SetOnClose(onClose ConnectionOnClose)
	SetOnRecvMessage(onRecvMessage ConnectionOnRecvMessage)
	Start() error
	Stop() error
}
