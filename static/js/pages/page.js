//生成table tr和page标签
function tablePage(e, title, d) {
    if (e.success) {
        $(e.data).each(function (i, o) {
            trs(i, o, title,d)
        })
        //    分页
        if (e.page !== undefined && e.page !== null) {
            $(e.page.pageData).each(function (i, o) {
                let a = (o === pageNum ? 'active' : '')
                let l = '<li class="page-item ' + a + '" onclick="pageLabel(\'' + o + '\')"><a class="page-link" href="#">' + o + '</a></li>'
                $("#table-page").append(l)
            })
        }
    } else {
        alter2(3, e.msg)
    }
}

//tr模板
function trs(i, data, title,d) {
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
            if("user" === d) {
                $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-warning btn-xs\" onclick=\"restPwd(\'" + data.uuid + "\')\">重置密码</button>")
            }
        } else {
            $(tr).find('td:last').text(data[e2])
        }
    })
}