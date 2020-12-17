var Toast;
$(function () {
    let t = sessionStorage.getItem("xueSeToken")
    if (!t) {
        window.location.href = window.origin + "/"
    }
    //弹窗全局设置
    Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
    });

    //----------------------------------------------------------------------

    //注销登录
    $("#logout").click(function () {
        sessionStorage.clear()
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

    getInfo().then(function () {
        return getNotice()
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
    return myAjax("/api/notice/notice", "GET", null, function (e) {
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
    return myAjax("/api/user/userInfo", "GET", null, function (e) {
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