package websocket

import (
	"net/http"
	"net/url"
)

// Client websocket client
type Client struct {
	ID    string
	Conn  Connection
	Data  interface{}
	Group string
}

// Command websocket command
type Command struct {
	Action string `json:"a"`
	Params string `json:"p"`
}

// CommandErrorParams websocket command error params
type CommandErrorParams struct {
	Message string `json:"msg"`
}

// Events websocket events
type Events struct {
	OnConnect     OnConnect
	OnRecvMessage OnRecvMessage
	OnClose       OnClose
}

// Message  message
type Message struct {
	Type    int
	Message []byte
}

// Settings websocket settings
type Settings struct {
	Port   int
	Events EventsMap
}

// BroadcastFilter websocket broadcast filter func
type BroadcastFilter func(client *Client) bool

// ChannelsMap websocket channels map
type ChannelsMap map[string]chan interface{}

// CommandProcess Command process func
type CommandProcess func(command *Command, client *Client) error

// CommandProcessMap Command process map
type CommandProcessMap map[string]CommandProcess

// ConnectionOnRecvMessage  message on receive func
type ConnectionOnRecvMessage func(message Message, conn Connection) error

// ConnectionOnClose  message on close
type ConnectionOnClose func(conn Connection)

// EventsMap websocket events map
type EventsMap map[string]Events

// HandlersMap http handlers map
type HandlersMap map[string]http.HandlerFunc

// OnClose websocket on close func
type OnClose func(client *Client)

// OnConnect websocket on connect func
type OnConnect func(urlQuery url.Values, client *Client) error

// OnRecvMessage websocket message on receive func
type OnRecvMessage func(message Message, client *Client) error
