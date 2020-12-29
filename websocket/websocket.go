package websocket

// Websocket Websocket
type Websocket interface {
	BroadcastMessage(message string) error
	BroadcastMessageByFilter(message string, filter BroadcastFilter) error
	BroadcastMessageByGroup(message, group string) error
	GetClientByID(id string) (*Client, error)
	GetClients() []*Client
	GetClientsByFilter(filter BroadcastFilter) []*Client
	GetClientsByGroup(group string) []*Client
	GetHandlerMap() HandlersMap
	RemoveClientByFilter(filter BroadcastFilter) error
	RemoveClientByGroup(group string) error
	RemoveClientByID(id string) error
	RemoveClients() error
	SendMessageByClientID(clientID string, message string) error
	Start() error
	Stop() error
}
