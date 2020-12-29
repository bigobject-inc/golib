package websocket

import (
	"errors"
)

// ErrAlreadyClosed Websocket connection is already closed
var ErrAlreadyClosed error = errors.New("Websocket connection is already closed")

// ErrAlreadyRunning websocket already running
var ErrAlreadyRunning error = errors.New("Websocket is already running")

// ErrClientEventNotFound websocket event not found
var ErrClientEventNotFound error = errors.New("Websocket event not found")

// ErrClientNotFound websocket client not found
var ErrClientNotFound error = errors.New("Websocket client not found")

// ErrNotRunning websocket is not running
var ErrNotRunning error = errors.New("Websocket is not running")
