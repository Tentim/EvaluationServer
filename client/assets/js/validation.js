
//  ----------------  验证

//$.growl.error({title: "错误标题", message: "错误消息内容!" });
//$.growl.notice({title: "提醒标题", message: "提醒消息内容!" });
//$.growl.warning({title: "警告标题", message: "警告消息内容!" });

    
//登录验证
function login_validation() {
    console.log("进行登录验证检测")
    var passwd = document.forms["loginForm"]["Password"].value;
    var name = document.forms["loginForm"]["Username"].value;
    var user = new proto.pb.User()
    user.setPassword(passwd)
    user.setUsername(name)
    var msg = new proto.pb.ClientMessage()
    msg.setOrder(proto.pb.ClientOrder.CLIENORDER_LOGIN) 
    msg.setUser(user)
    console.log(user.toObject())

    //序列化
    var S = msg.serializeBinary()
    ws.send(S)
    return false
}

//注册验证
function signup_validation() {
    var passwd = document.forms["signupForm"]["Password"].value;
    var reppasswd = document.forms["signupForm"]["RepPassword"].value;
    var name = document.forms["signupForm"]["Username"].value;
    if(passwd != reppasswd){
        console.log(passwd, reppasswd)
        $.growl.error({ title: "两次密码不正确", message: "请重新输入" });
        return false
    }
    var user = new proto.pb.User()
    user.setPassword(passwd)
    user.setUsername(name)
    var msg = new proto.pb.ClientMessage()
    msg.setOrder(proto.pb.ClientOrder.CLIENORDER_SIGNUP)
    msg.setUser(user)
    console.log(msg.toObject())

    //序列化
    var S = msg.serializeBinary()
    ws.send(S)
    return false
}

//确认密码验证
function signup_passwd_Rep() {
    var passwd = document.forms["signupForm"]["Password"].value;
    var reppasswd = document.forms["signupForm"]["RepPassword"].value;
    console.log(passwd, reppasswd)
    if(passwd != reppasswd){
        $.growl.warning({ title: "两次密码不正确", message: "请重新输入" });
    }
}

//管理员验证
function admin_validation() {
    var passwd = document.forms["adminForm"]["Password"].value;
    var name = document.forms["adminForm"]["Username"].value;
    console.log(passwd, reppasswd)
}

//进行检测
function test_t() {
    console.log("测试")
    //window.location.href = "http://www.baidu.com";
    //window.open("http://www.baidu.com");
}
