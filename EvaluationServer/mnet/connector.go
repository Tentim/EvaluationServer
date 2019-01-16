package mnet

import "golang.org/x/net/websocket"

//Connector 客户端结构体
type Connector struct {
	ws websocket.Conn
}
