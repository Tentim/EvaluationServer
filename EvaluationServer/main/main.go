package main

import (
	"EvaluationServer/bank"
	"EvaluationServer/mnet"
	"EvaluationServer/msql"
	"fmt"
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
	http.Handle("/ans", websocket.Handler(AnswerServerHandle))

	//打印日志
	log.Println("服务器启动成功！")

	//TODO
	//设置时间
	setTime(0, 1, 0)

	var code int

	//监听端口
	go func() {
		QIDs = bank.GenerateRandomNumber(1, 30, QCount) //选题
		if err := http.ListenAndServe(":4545", nil); err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	for {
		fmt.Printf("请输入指令: ")
		if _, err := fmt.Scan(&code); err != nil || code < 0 || code > 2 {
			fmt.Println("指令错误: ")
			continue
		}

		switch code {
		case 0:
			{
				go timeStart()
				fmt.Println("开始等待倒计时")
			}
		case 1: //初始化更新题库
			{
				go bank.UpDataQues()
				fmt.Println("数据库已更新")
			}
		case 2:
			{
				go func() {
					QIDs = bank.GenerateRandomNumber(1, 30, QCount)
					for _, v := range QIDs {
						if ques, ok := msql.GetQuesByID(v); ok {
							log.Println(v, ":", ques)
						} else {
							log.Println("题库读取出错")
						}
					}
				}()
			}
		}
	}
}
