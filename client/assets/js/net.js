//声明 websocket 变量
var ws 

if( "WebSocket" in window ){

    //新建连接
    ws = new WebSocket("ws://127.0.0.1:4545");
    
    //设置连接事件
    ws.onopen = function(){
        console.log("connect websocket success")
    }

    //接受消息
    ws.onmessage = function(evt){
        console.log("收到服务器消息")
        var reader = new FileReader(); 
        reader.readAsArrayBuffer(evt.data);//读取数据
        reader.onload = function (e) {
            servermsg = proto.pb.ServerMessage.deserializeBinary( reader.result ) //反序列化
            switch(servermsg.getOrder()){
                case proto.pb.ServerOrder.SERERORDER_LOGIN:{
                    var loginmsg = servermsg.getLogin()
                    console.log(loginmsg.toObject())
                    if( loginmsg.getIstrue()){
                        window.location.href = "http://127.0.0.1/04/assets/html/answer.html";
                    }else{
                        alert("用户名或密码错误")
                    }      
                };break;
            }
        }
    }

    //设置断开连接事件
    ws.onclose = function(){
        console.log("Close WebSocket")
    }

}else{
    alert("No WebSocket")
}