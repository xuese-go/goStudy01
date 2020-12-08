var t;
var Toast;
$(function () {
    t = window.localStorage.getItem("xueSeToken")
    if (!t) {
        window.location.href = window.origin + "/"
    }
    //弹窗全局设置
    Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
    });
//    全局设置
    $.ajaxSetup({
        timeout: 2000,
        dataType: "json",
        beforeSend: function (request) {
            request.setRequestHeader("xueSeToken", t);
        },
        error: function (err) {
            console.log(err)
            if (err.status === 200) {
                if (!err.responseJSON.success) {
                    if (err.responseJSON.data === 'logout') {
                        alter2(4, err.responseJSON.msg)
                        setTimeout(function () {
                            window.location.href = window.origin + "/"
                        }, 3000);
                    } else if (err.responseJSON.data === '!admin') {
                        alter2(4, err.responseJSON.msg)
                    }
                } else {
                    console.log(err.responseJSON)
                }
            } else if (err.status === 404) {
                alter2(4, "资源不存在")
            } else {
                alter2(4, "资源错误")
            }
        }
    });
    //----------------------------------------------------------------------

    //获取消息
    getNotice()
    //获取当前登录用户信息
    getInfo()

    //注销登录
    $("#logout").click(function () {
        window.localStorage.clear()
        window.location.href = window.origin + "/"
    })

//    菜单点击跳转
    $(".nav-link").on("click", function () {
        let id = $(this).attr("data-id")
        if (id) {
            $.ajax("/page" + id, {
                type: "GET",
                dataType: 'html'
            }).done(function (e) {
                $("#content").html(e)
            })
        }
    })
})

//弹窗提示
function alter2(icon, title) {
    switch (icon) {
        case 1:
            icon = 'success'
            break
        case 2:
            icon = 'info'
            break
        case 3:
            icon = 'warning'
            break
        case 4:
            icon = 'error'
            break
    }
    Toast.fire({
        icon: icon,
        title: title,
        showConfirmButton: false,
        timer: 2000
    })
}

//弹窗提示
function alter2IsOk(butText) {
    return Swal.fire({
        icon: 'warning',
        showConfirmButton: true,
        confirmButtonColor: '#3085d6',
        confirmButtonText: butText,
        showCancelButton: true,
        cancelButtonColor: '#d33',
        cancelButtonText: "取消"
    })
}

//获取所有通知
function getNotice() {
    $.ajax("/api/notice/notice", {
        type: "GET",
        dataType: 'json'
    }).done(function (e) {
        if (e.success) {
            $("#notice-num").text(e.data.length)
            $(e.data).each(function (i, obj) {
                if (i <= 5) {
                    let h = '<a href="#" class="dropdown-item">' +
                        '<i class="fas fa-envelope mr-2"></i>' +
                        obj + '<span class="float-right text-muted text-sm">3 mins</span>' +
                        '</a>' +
                        '<div class="dropdown-divider"></div>'
                    $("#notice-item").append(h)
                }
            })
        }
    })
}

function getInfo() {
    $.ajax("/api/user/user", {
        type: "GET",
        dataType: 'json'
    }).done(function (e) {
        if (e.success) {
            if (e.data.role === 2) {
                $(".admin").show()
            }
            $("#info").text(e.data.account)
            if (e.data.portrait) {
                $("#home-user-img").attr("src", "/file/" + e.data.portrait)
            }
        } else {
            alter2(4, "个人信息获取失败")
        }
    })
}