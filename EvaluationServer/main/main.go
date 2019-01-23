package main

import (
	"EvaluationServer/mnet"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

//Hub 集线器
var Hub *mnet.Hub

// 主程序
func main() {

	//设置日志格式
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds)

	//创建集线器
	Hub = mnet.NewHub()

	//设置路由
	http.Handle("/", websocket.Handler(ServerHandle))

	//打印日志
	log.Println("服务器启动成功！")

	//监听端口
	if err := http.ListenAndServe(":4545", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}
