
//登录验证
function login_validation() {
    console.log("进行登录验证检测")
    var passwd = document.forms["loginForm"]["Password"].value;
    var name = document.forms["loginForm"]["Username"].value;
    var user = new proto.pb.User()
    user.setPassword(passwd)
    user.setUsername(name)
    var msg = new proto.pb.ClientMessage()
    msg.setOrder = proto.pb.ClientOrder.CLIENORDER_LOGIN
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
    var name = document.forms["signupForm"]["Username"].value;
}

//管理员验证
function admin_validation() {
    var passwd = document.forms["adminForm"]["Password"].value;
    var name = document.forms["adminForm"]["Username"].value;
}

//进行检测
function test() {
    console.log("检测")
    //window.location.href = "http://www.baidu.com";
    window.open("http://www.baidu.com");
}


