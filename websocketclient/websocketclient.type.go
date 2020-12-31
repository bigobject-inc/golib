package websocketclient

import (
	"net/url"
)

// Settings websocket client settings
type Settings struct {
	URL           url.URL
	OnRecvMessage OnRecvMessage
}

// OnRecvMessage websocket client receive message
type OnRecvMessage func(messageType int, message []byte) error
