package golib

import (
	"github.com/bigobject-inc/golib/dbquery"
	"github.com/bigobject-inc/golib/irisweb"
	"github.com/bigobject-inc/golib/logger"
	"github.com/bigobject-inc/golib/miscellaneous"
	"github.com/bigobject-inc/golib/websocket"
	"github.com/bigobject-inc/golib/websocketclient"
)

// GoLibrary golang library service
type GoLibrary interface {
	GetServMiscellaneous() miscellaneous.Miscellaneous
	GetVersion() string
	NewLogger(loggerFile, loggerLevel string) (logger.Logger, logger.Defer, error)
	NewRotationLogger(loggerFile, loggerLevel string) (logger.Logger, error)
	NewWeb(l logger.Logger, conf irisweb.Configure, rootPath string) (irisweb.Web, error)
	NewWebController(path, method, name, summary string) (irisweb.Controller, error)
	NewWebSocket(logger logger.Logger, settings websocket.Settings) (websocket.Websocket, error)
	NewWebSocketClient(logger logger.Logger, settings websocketclient.Settings) (websocketclient.WebsocketClient, error)
	NewPostgres(username, password, ip, port, database string) dbquery.Dbquery
	NewBO(username, password, ip, port, database string) dbquery.Dbquery
	NewBOQuery(username, password, ip, port, database string) dbquery.Boquery
}
