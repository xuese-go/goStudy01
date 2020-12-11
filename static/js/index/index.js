// 提交
$(function () {
    $("#login-form").on("submit", function (ev) {
        ev.preventDefault();
        $.ajax("/api/login/login", {
            type: "POST",
            dataType: 'json',
            data: $('#login-form').serialize()
        }).done(function (e) {
            if (e.success) {
                sessionStorage.setItem("xueSeToken", e.data);
                if (sessionStorage.getItem("xueSeToken")) {
                    window.location.href = window.origin + "/page/home"
                }
            } else {
                $("#msg").text(e.msg)
            }
        }).fail(function (err) {
            $("#msg").text("账号或密码错误")
        })
    })
})