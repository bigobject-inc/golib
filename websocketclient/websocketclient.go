package websocketclient

import (
	"context"
)

// WebsocketClient Websocket client
type WebsocketClient interface {
	IsRunning() bool
	SendText(message []byte) error
	Start(ctx context.Context) error
	Stop() error
}
