var pageNum = 1
var pageSize = 8
$(function () {
    //分页
    page()
    //查询
    $("#form-search").bind("click", function () {
        page()
    })
    //新增
    $("#form-save").on("submit", function (ev) {
        ev.preventDefault();
        return myAjax("/api/user/user", "POST", $('#form-save').serialize(), function (e) {
            if (e.success) {
                alter2(1, "成功")
                $("#btn-to-save").click()
                $('#form-save')[0].reset()
                page()
            } else {
                alter2(3, e.msg)
            }
        })
    })
//    修改
    $("#form-update").on("submit", function (ev) {
        ev.preventDefault();
        return myAjax("/api/user/user/" + $("#uuid2").val(), "PUT", $('#form-update').serialize(), function (e) {
            if (e.success) {
                alter2(1, "成功")
                $("#btn-to-update").click()
                $('#form-update')[0].reset()
                page()
            } else {
                alter2(3, e.msg)
            }
        })
    })
})

//分页标签
function pageLabel(o) {
    pageNum = o
    page()
}

//分页查询
function page() {
    $("#table-content").find("tr").remove()
    $("#table-page").find("li").remove()
    return myAjax("/api/user/user", "GET", {
        "pageNum": pageNum,
        "pageSize": pageSize,
        "account": $("#table_search").val()
    }, function (e) {
        let title = ["account", "role", "state", "btn"]
        tablePage(e, title)
    })
}

//tr模板
function trs(i, data, title) {
    $("#table-content").append("<tr>")
    let tr = $('#table-content').find('tr:last');

    $(tr).append("<td>")
    $(tr).find('td:last').text(i + 1)

    $(title).each(function (i2, e2) {
        $(tr).append("<td>")
        if (e2 === "role") {
            let r = (data[e2] === 1 ? '<span class="badge bg-primary">普通</span>' : '<span class="badge bg-danger">管理员</span>')
            $(tr).find('td:last').html(r)
        } else if (e2 === "state") {
            let r = (data[e2] === 1 ? '<span class="badge bg-success">正常</span>' : '<span class="badge bg-secondary">停用</span>')
            $(tr).find('td:last').html(r)
        } else if (e2 === "btn") {
            $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-danger btn-xs\" onclick=\"del(\'" + data.uuid + "\')\">删除</button>")
            $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-warning btn-xs\" onclick=\"one(\'" + data.uuid + "\')\" data-toggle=\"modal\" data-target=\"#modal-update\">修改/查看</button>")
            $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-warning btn-xs\" onclick=\"restPwd(\'" + data.uuid + "\')\">重置密码</button>")
        } else {
            $(tr).find('td:last').text(data[e2])
        }
    })
}

//删除
function del(o) {
    alter2IsOk("是否确定删除？").then(function (e) {
        if (e.value) {
            return myAjax("/api/user/user/" + o, "DELETE", null, function (e) {
                if (e.success) {
                    page()
                } else {
                    alter2(4, e.msg)
                }
            })
        }
    })
}

//根据id获取
function one(e) {
    return myAjax("/api/user/user/" + e, "GET", null, function (e) {
        if (e.success) {
            $("#uuid2").val(e.data.uuid)
            $("#exampleInputEmail12").val(e.data.account)
            $("#role2").val(e.data.role)
            $("#state2").val(e.data.state)
        } else {
            alter2(3, e.msg)
        }
    })
}

//根据id重置密码
function restPwd(e) {
    return myAjax("/api/user/rest/pwd/" + e, "GET", null, function (e) {
        if (e.success) {
            alter2(1, "重置成功")
        } else {
            alter2(3, e.msg)
        }
    })
}