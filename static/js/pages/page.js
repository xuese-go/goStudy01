//生成table tr和page标签
function tablePage(e, title) {
    if (e.success) {
        $(e.data).each(function (i, o) {
            trs(i, o, title)
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
// function trs(i, e, data) {
//     $("#table-content").append("<tr>")
//     let tr = $('#table-content').find('tr:last');
//
//     $(tr).append("<td>")
//     $(tr).find('td:last').text(i + 1)
//
//     $(data).each(function (i2, e2) {
//         $(tr).append("<td>")
//         $.each(e, function (i3) {
//             if (i3 === e2) {//键
//                 $(tr).find('td:last').text(e[i3])//获取对应的value
//             }
//         });
//     })
//
//     $(tr).append("<td>")
//     $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-danger btn-xs\" onclick=\"del(\'" + e.uuid + "\')\">删除</button>")
//     $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-warning btn-xs\" onclick=\"one(\'" + e.uuid + "\')\" data-toggle=\"modal\" data-target=\"#modal-update\">修改/查看</button>")
// }