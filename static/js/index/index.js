// 提交
$(function(){
    $("#login-btn").on("click",function(){
        $.ajax("/api/login/login",{
            type: "POST",
            dataType: 'json',
            data: $('#login-form').serialize()
        }).done(function(e){
            if(e.success){
                window.localStorage.setItem("xueSeToken",e.data);
                if(window.localStorage.getItem("xueSeToken")) {
                    window.location.href = window.origin + "/page/home"
                }
            }else{
                $("#msg").text(e.msg)
            }
        }).fail(function(err){
            $("#msg").text("系统错误")
        })
    })
})