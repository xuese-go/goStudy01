// 提交
$(function(){
    $("#login-btn").on("click",function(){
        $.ajax("/api/login/login",{
            type: "POST",
            dataType: 'json',
            data: $('#login-form').serialize()
        }).done(function(e){
            console.log()
            // window.location.href="/"
        }).fail(function(err){
            console.log(err.responseJSON.msg)
            console.log("err")
        })
    })
})