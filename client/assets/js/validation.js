
//登录验证
function login_validation() {
    console.log("进行登录验证检测")
    var passwd = document.forms["loginForm"]["Password"].value;
    var name = document.forms["loginForm"]["Username"].value;
    ws.send(name)
    ws.send(passwd)
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
}


