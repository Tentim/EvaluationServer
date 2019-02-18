var wait = function(){

    var fun = function (o) {
        o.show("正在进入答题界面。。。");
        setTimeout(o.hide, 3000);
        setTimeout(function () {window.location.href = "http://127.0.0.1/Eva/assets/html/answer.html"}, 3000);
    }
    
    seajs.use("seajs-waiting", fun)
} 
