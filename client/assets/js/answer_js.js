
/**
 * 答题
 * Created by pjj on 2017/2/4.
 */
window.onload=function clear_all(){
    var a = document.getElementsByTagName("input");
    for(var i = 0;i < a.length;i++){
        if(a[i].type == "radio"){
            a[i].checked = false;
        }
    }
};

function score_check(){
    var sum= 0;
    var obj  = document.getElementsByTagName('input');
    for(var i=0;i<obj.length;i++){
        if(obj[i].checked==true){
            if(obj[i].value=='*') sum++;
        }
    }
    var test=document.getElementById('1'),
        score=document.getElementById('2');
    test.style.display='none';
    score.style.display='block';

    var q=document.getElementById('q'),
        s=document.getElementById('s');
    q.innerHTML= sum ;
    s.innerHTML= sum*20;
}

function return_test(){
    var test=document.getElementById('1'),
        score=document.getElementById('2');
    test.style.display='block';
    score.style.display='none';
}

$(function () {
    //发送获取题库指令
    var msg = new proto.pb.ClientMessage()
    msg.setOrder(proto.pb.ClientOrder.CLIENORDER_GET_QUESTION)

    //序列化
    var S = msg.serializeBinary()
    if(ws.readyState == 1){
        ws.send(S)
    }
    else{
        ws.onopen = function(){ws.send(S)} 
    }
})

