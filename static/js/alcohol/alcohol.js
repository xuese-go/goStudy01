var pageNum = 1
var pageSize = 8
$(function () {
//分页
    page()
//    加载所有品牌
    loadBrand()
//    加载所有系列
    loadSeries()
//查询
    $("#form-search").bind("click", function () {
        page()
    })
//新增
    $("#form-save").on("submit", function (ev) {
        ev.preventDefault();
        $.ajax("/api/alcohol/alcohol", {
            type: "POST",
            dataType: 'json',
            data: $('#form-save').serialize()
        }).done(function (e) {
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
        $.ajax("/api/alcohol/alcohol/" + $("#uuid2").val(), {
            type: "PUT",
            dataType: 'json',
            data: $('#form-update').serialize()
        }).done(function (e) {
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
    $.ajax("/api/alcohol/alcohols", {
        type: "GET",
        dataType: 'json',
        data: {
            "pageNum": pageNum,
            "pageSize": pageSize,
            "name": $("#table_search").val()
        }
    }).done(function (e) {
        if (e.success) {
            $(e.data).each(function (i, o) {
                $("#table-content").append(trs(i, o))
            })
//    分页
            if (e.page !== undefined && e.page !== null) {
                $(e.page.pageData).each(function (i, o) {
                    let a = (o === pageNum ? 'active' : '')
                    let l = '<li class="page-item ' + a + '" onclick="pageLabel(\'' + o + '\')"><a class="page-link" href="#">' + o + '</a></li>'
                    $("#table-page").append(l)
                })
            }
        }
    })
}

//tr模板
function trs(i, e) {
    $("#table-content").append("<tr>")
    let tr = $('#table-content').find('tr:last');

    $(tr).append("<td>")
    $(tr).find('td:last').text(i + 1)

    $(tr).append("<td>")
    $(tr).find('td:last').text(e.name)
    $(tr).append("<td>")
    $(tr).find('td:last').text(e.brandId)
    $(tr).append("<td>")
    $(tr).find('td:last').text(e.seriesId)
    $(tr).append("<td>")
    $(tr).find('td:last').text(e.concentration)

    $(tr).append("<td>")
    $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-danger btn-xs\" onclick=\"del(\'" + e.uuid + "\')\">删除</button>")
    $(tr).find('td:last').append("&nbsp;&nbsp;<button type=\"button\" class=\"btn btn-warning btn-xs\" onclick=\"one(\'" + e.uuid + "\')\" data-toggle=\"modal\" data-target=\"#modal-update\">修改/查看</button>")
}

//删除
function del(o) {
    alter2IsOk("是否确定删除？").then(function (e) {
        if (e.value) {
            $.ajax("/api/alcohol/alcohol/" + o, {
                type: "DELETE",
                dataType: 'json'
            }).done(function (e) {
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
    $.ajax("/api/alcohol/alcohol/" + e, {
        type: "GET",
        dataType: 'json'
    }).done(function (e) {
        if (e.success) {
            $("#uuid2").val(e.data.uuid)
            $("#name2").val(e.data.name)
            $("#brandId2").val(e.data.brandId)
            $("#seriesId2").val(e.data.seriesId)
            $("#concentration2").val(e.data.concentration)
        } else {
            alter2(3, e.msg)
        }
    })
}

//加载所有品牌
function loadBrand() {
    $("#brandId").find("option").remove()
    $("#brandId2").find("option").remove()
    $.ajax("/api/brand/brands", {
        type: "GET",
        dataType: 'json',
        data: {
            "pageNum": pageNum,
            "pageSize": pageSize,
        }
    }).done(function (e) {
        if (e.success) {
            $(e.data).each(function (i, o) {
                $("#brandId").append("<option value='" + o.uuid + "'>" + o.name + "</option>>")
                $("#brandId2").append("<option value='" + o.uuid + "'>" + o.name + "</option>>")
            })
        }
    })
}
//加载所有系列
function loadSeries() {
    $("#seriesId").find("option").remove()
    $("#seriesId2").find("option").remove()
    $.ajax("/api/series/seriess", {
        type: "GET",
        dataType: 'json',
        data: {
            "pageNum": pageNum,
            "pageSize": pageSize,
        }
    }).done(function (e) {
        if (e.success) {
            $(e.data).each(function (i, o) {
                $("#seriesId").append("<option value='" + o.uuid + "'>" + o.name + "</option>>")
                $("#seriesId2").append("<option value='" + o.uuid + "'>" + o.name + "</option>>")
            })
        }
    })
}