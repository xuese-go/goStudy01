$(function () {

    //个人信息
    getInfo()

    $(".img-circle2").bind("click", function (ev) {
        $("#file").click()
    })
    $("#file").bind("change", function (ev) {
        let formData = new FormData();
        // HTML 文件类型input，由用户选择
        formData.append("file", $(this)[0].files[0]);
        $.ajax({
            url: '/api/user/file',
            dataType: 'json',
            type: 'PUT',
            async: false,
            data: formData,
            processData: false, // 使数据不做处理
            contentType: false, // 不要设置Content-Type请求头
            success: function (e) {
                if (e.success) {
                    $("#user-image").attr("src", window.origin + "/file/" + e.data)
                }
            },
            error: function (response) {
                console.log(response);
            }
        });
    })
})

function getInfo() {
    $.ajax("/api/user/user", {
        type: "GET",
        dataType: 'json'
    }).done(function (e) {
        if (e.success) {
            if (e.data.portrait) {
                $("#user-image").attr("src", "/file/" + e.data.portrait)
            }
        } else {
            alter2(4, "个人信息获取失败")
        }
    }).fail(function (err) {

    })
}