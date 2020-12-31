package websocketclient

import (
	"context"
)

// WebsocketClient Websocket client
type WebsocketClient interface {
	IsRunning() bool
	Start(ctx context.Context) error
	Stop() error
}
