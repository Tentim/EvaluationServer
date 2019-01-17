package main

import (
	"EvaluationServer/mnet"
	"sync"

	"golang.org/x/net/websocket"
)

//从客户端读取消息
func readFromClientConn(ws *websocket.Conn, wg *sync.WaitGroup, connid int64) {
	defer wg.Done()
}

// ServerHandle 服务器处理函数
func ServerHandle(ws *websocket.Conn) {

	//生成客户端id
	connid := mnet.CreateSessionID()

	//添加到集线器
	Hub.Add(connid, ws)

	//从集线器中删除该
	defer Hub.ClientClose(connid)

	//读取读取客户端消息
	var wg sync.WaitGroup
	for index := 0; index < 5; index++ {
		wg.Add(1)
		go readFromClientConn(ws, &wg, connid)
	}

	wg.Wait()
}
