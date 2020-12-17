//    全局设置
$.ajaxSetup({
    timeout: 2000,
    dataType: "json",
    beforeSend: function (request) {
        request.setRequestHeader("xueSeToken", sessionStorage.getItem("xueSeToken"));
        $("#loading").show()
        $("#content").hide()
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
            for (let p in err) {
                if (err.hasOwnProperty(p)) {
                    if ("responseJSON" === p) {
                        for (let p1 in err[p]) {
                            if (err[p].hasOwnProperty(p1)) {
                                if ("data" === p1) {
                                    if(err[p][p1] === "logout"){
                                        $('div').remove()
                                        $('body').append("<span>"+err.responseJSON.msg+"3秒后跳转登录页</span>")
                                        setTimeout(function () {
                                            window.location.href = window.origin + "/"
                                        }, 3000);
                                    }else if(err[p][p1] === "!admin"){
                                        alter2(4, err.responseJSON.msg)
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    },
    complete: function () {
        setTimeout(function () {
            $("#loading").hide()
            $("#content").show()
            $(":submit").removeClass("disabled")
        }, 500);
    }
});

function myAjax(url, type, data, callback) {
    return new Promise(function (resolve, reject) {
        $.ajax(url, {
            type: type,
            dataType: 'json',
            data: data
        }).done(function (e) {
            callback(e)
            resolve()
        }).fail(function () {
            reject();
        })
    })
}