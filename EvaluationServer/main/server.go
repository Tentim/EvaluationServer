package main

import (
	"EvaluationServer/mnet"
	"EvaluationServer/pb"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/websocket"
)

//从客户端读取消息
func readFromClientConn(ws *websocket.Conn, wg *sync.WaitGroup, connid int64) {
	defer wg.Done()
	for {
		//读取数据
		var pData []byte
		if err := websocket.Message.Receive(ws, &pData); err != nil {
			return
		}

		//反序列化
		var msg = &pb.User{}
		if err := proto.Unmarshal(pData, msg); err != nil {
			log.Println("反序列化错误")
			return
		}

		log.Println(msg)

	}
}

//客户端退出处理函数
func delect(id int64) {
	Hub.ClientClose(id)
	log.Println(id, "已退出")
	log.Println("hub: {", Hub.GetAll(), "}")
}

// ServerHandle 服务器处理函数
func ServerHandle(ws *websocket.Conn) {

	//生成客户端id
	connid := mnet.CreateSessionID()

	//添加到集线器
	Hub.Add(connid, ws)

	//从集线器中删除该
	defer delect(connid)

	log.Println(connid, "连接到服务器")
	log.Println("hub: {", Hub.GetAll(), "}")

	//读取读取客户端消息
	var wg sync.WaitGroup
	for index := 0; index < 5; index++ {
		wg.Add(1)
		go readFromClientConn(ws, &wg, connid)
	}

	wg.Wait()
}
