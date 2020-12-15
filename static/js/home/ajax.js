//    全局设置
$.ajaxSetup({
    timeout: 2000,
    dataType: "json",
    beforeSend: function (request) {
        request.setRequestHeader("xueSeToken", sessionStorage.getItem("xueSeToken"));
        $(":submit").addClass("disabled")
    },
    success: function (response, status, xhr) {
        //响应头部
        let to = xhr.getResponseHeader("token")
        if (to != null) {
            sessionStorage.setItem("xueSeToken", to);
        }
    },
    error: function (err) {
        if (err.status === 404) {
            alter2(4, "资源不存在")
        } else {
            if (err.responseJSON.success !== undefined && !err.responseJSON.success) {
                if (err.responseJSON.data === 'logout') {
                    $('div').remove()
                    $('body').append("<span>"+err.responseJSON.msg+"3秒后跳转登录页</span>")
                    setTimeout(function () {
                        window.location.href = window.origin + "/"
                    }, 3000);
                } else if (err.responseJSON.data === '!admin') {
                    alter2(4, err.responseJSON.msg)
                } else {
                    alter2(4, err.responseJSON.msg)
                }
            } else {
                alter2(4, "资源错误")
            }
        }
    },
    complete: function () {
        $(":submit").removeClass("disabled")
    }
});

function myAjax(url, type, data,callback) {
    var p = new Promise(function (resolve, reject) {
        $.ajax(url, {
            type: type,
            dataType: 'json',
            data: data
        }).done(function (e) {
            callback(e)
            resolve()
        }).fail(function(){
            reject();
        })
    })
    return p
}