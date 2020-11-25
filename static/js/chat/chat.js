$(function () {
//  获取所有用户
    findUsers()
//    搜索用户
    $("#chat-user-search").on("click", function (ev) {
        if ($("#chat-user-search-input").val() !== "") {
            $("#chat-user").find(".chat-user").each(function (i, e) {
                if ($(e).text().indexOf($("#chat-user-search-input").val()) > -1) {
                    $(e).show()
                } else {
                    $(e).hide()
                }
            })
        } else {
            $("#chat-user").find(".chat-user").show()
        }
    })
})

function findUsers() {
    $("#chat-user").find(".chat-user").remove()
    $.ajax("/api/user/users", {
        type: "GET",
        dataType: 'json',
        data: {
            "pageNum": 1,
            "pageSize": 100,
            "account": ""
        }
    }).done(function (e) {
        if (e.success) {
            $(e.data).each(function (i, o) {
                $("#chat-user").append(trs(i, o))
            })
        } else {

        }
    }).fail(function (err) {

    })
}

//tr模板
function trs(i, e) {
    return '<div class="row col-12 chat-user" onclick="toChat(\'' + e.uuid + '\',\'' + e.account + '\')">' +
        '<img class="direct-chat-img" src="' + (e.portrait === "" ? "/static/img/user1-128x128.jpg" : ("/file/" + e.portrait)) + '" >' +
        e.account +
        '</div>'
}

//获取与当前点击人的聊天记录
function toChat(a, b) {
    $("#toChat").text(b)
    $("#chat-msg").find("div").remove()
}