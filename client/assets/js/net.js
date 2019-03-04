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
            console.log(servermsg.toObject())
            switch(servermsg.getOrder()){
                case proto.pb.ServerOrder.SERERORDER_LOGIN:{
                    var loginmsg = servermsg.getLogin()
                    if( loginmsg.getIstrue()){
                        //window.location.href = "http://127.0.0.1/Eva/assets/html/time.html";
                        window.location.href = "time.html";
                    }else{
                        alert("用户名或密码错误")
                    }
                };break;
                case proto.pb.ServerOrder.SERERORDER_SIGNUP:{
                    var signup = servermsg.getSignup()
                    if( signup.getIstrue()){
                        window.location.href = "login.html";
                    }else{
                        alert("注册失败")
                    } 
                };break;
                case proto.pb.ServerOrder.SERERORDER_SEND_WAITTIME:{
                    var wait = servermsg.getWait()
                    var start = wait.getStart()
                    var waittime = wait.getTime()
                    console.log(waittime.toObject())
                    console.log(start)
                    setTimeCountDown(waittime.getHour(),waittime.getMinute(), waittime.getSecond())
                    if(start){
                        startCountDown()
                    }
                };break;
                case proto.pb.ServerOrder.SERERORDER_SEND_QUESTION:{
                    console.log("更新题库")
                    var quess = servermsg.getQuess()
                    var qcoust = quess.getNum()
                    var ques = quess.getQuesList()
                    for(i=0; i<qcoust; i++){
                        console.log(i)
                        var tpl = template(document.getElementById('tpl').innerHTML);
                        var html = tpl({N:i, Q: ques[i].getQuestion(), A: ques[i].getA(), B: ques[i].getB(), C: ques[i].getC(), D: ques[i].getD()});
                        var T = document.createElement("div")
                        T.innerHTML = html
                        document.getElementById('ques').appendChild(T);
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