//声明 websocket 变量
var ws 

if( "WebSocket" in window ){

    //新建连接
    ws = new WebSocket("ws://127.0.0.1:4545");
    
    //设置连接事件
    ws.onopen = function(){
        alert("OK WebSocket")
    }

    //设置断开连接事件
    ws.onclose = function(){
        alert("Close WebSocket")
    }

}else{
    alert("No WebSocket")
}