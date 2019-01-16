package main

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
)

// ServerHandle 服务器处理函数
func ServerHandle(ws *websocket.Conn) {
	log.Println(Hub)

	//添加到集线器
	Hub.Add(1, ws)

	for {
		log.Println(Hub)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
