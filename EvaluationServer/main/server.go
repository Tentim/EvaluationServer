package main

import (
	"EvaluationServer/bank"
	"EvaluationServer/mnet"
	"EvaluationServer/msql"
	"EvaluationServer/pb"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/websocket"
)

//注册处理函数 ClientOrder_CLIENORDER_SIGNUP 1
func signup(ws *websocket.Conn, user *pb.User) {
	log.Println("开始注册")

	//注册
	var DU msql.UserData
	DU.Username = user.GetUsername()
	DU.Password = user.GetPassword()
	log.Println(DU)
	res := msql.InsertUser(DU)

	//设置服务器消息
	signupMsg := &pb.SIGNUP{}
	signupMsg.Istrue = res
	serverMsg := &pb.ServerMessage{}
	serverMsg.Order = pb.ServerOrder_SERERORDER_SIGNUP
	serverMsg.Signup = signupMsg

	//序列化
	pData, err := proto.Marshal(serverMsg)
	if err != nil {
		log.Println("序列化错误")
	}

	//发送数据
	if err = websocket.Message.Send(ws, pData); err != nil {
		log.Println("发送数据错误")
	}

	log.Println("注册完成：", res)
}

//登录处理函数 ClientOrder_CLIENORDER_LOGIN 0
func login(ws *websocket.Conn, user *pb.User) {
	log.Println("开始登录验证")

	//密码验证
	res := msql.IsPasswdTrueByUsername(user.GetUsername(), user.GetPassword())

	//设置服务器消息
	loginMsg := &pb.Login{}
	loginMsg.Istrue = res
	serverMsg := &pb.ServerMessage{}
	serverMsg.Order = pb.ServerOrder_SERERORDER_LOGIN
	serverMsg.Login = loginMsg

	//序列化
	pData, err := proto.Marshal(serverMsg)
	if err != nil {
		log.Println("序列化错误")
	}

	//发送数据
	if err = websocket.Message.Send(ws, pData); err != nil {
		log.Println("发送数据错误")
	}

	log.Println("验证完成：", res)
}

func sendTime(ws *websocket.Conn) {
	//准备协议
	serverMsg := &pb.ServerMessage{}
	serverMsg.Order = pb.ServerOrder_SERERORDER_SEND_WAITTIME
	serverMsg.Time = waittime

	//序列化
	pData, err := proto.Marshal(serverMsg)
	if err != nil {
		log.Println("序列化错误")
	}

	//发送数据
	if err = websocket.Message.Send(ws, pData); err != nil {
		log.Println("发送数据错误")
	}

	log.Println("时间校准完成")

	go func() {
		nums := bank.GenerateRandomNumber(1, 30, 10)
		for _, v := range nums {
			if ques, ok := msql.GetQuesByID(v); ok {
				log.Println(v, ":", ques)
			} else {
				log.Println("题库读取出错")
			}
		}
	}()
}

//从客户端读取消息
func readFromClientConn(ws *websocket.Conn, wg *sync.WaitGroup, connid int64) {
	defer wg.Done()
	for {
		//读取数据
		var pData []byte
		if err := websocket.Message.Receive(ws, &pData); err != nil {
			return
		}

		log.Println("收到客户端消息")

		//反序列化
		var clientMsg = &pb.ClientMessage{}
		if err := proto.Unmarshal(pData, clientMsg); err != nil {
			log.Println("反序列化错误")
			return
		}

		//处理客户端事件
		switch clientMsg.GetOrder() {
		case pb.ClientOrder_CLIENORDER_LOGIN:
			{
				login(ws, clientMsg.GetUser())
			}
		case pb.ClientOrder_CLIENORDER_SIGNUP:
			{
				signup(ws, clientMsg.GetUser())
			}
		case pb.ClientOrder_CLIENORDER_GET_WAITTIME:
			{
				sendTime(ws)
			}
		} // end switch
	} // end for
} // end func

// ServerHandle 服务器处理函数
func ServerHandle(ws *websocket.Conn) {

	//生成客户端id
	connid := mnet.CreateSessionID()

	//添加到集线器
	//Hub.Add(connid, ws)

	//从集线器中删除该
	//defer Hub.ClientClose(connid)

	//log.Println(connid, "连接到服务器")
	//log.Println("hub: {", Hub.GetAll(), "}")

	//读取读取客户端消息
	var wg sync.WaitGroup
	for index := 0; index < 5; index++ {
		wg.Add(1)
		go readFromClientConn(ws, &wg, connid)
	}

	wg.Wait()
}
