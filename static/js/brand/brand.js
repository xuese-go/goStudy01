var pageNum = 1;
var pageSize = 8;
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
        return myAjax("/api/brand/brand", "POST", $('#form-save').serialize(), function (e) {
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
        return myAjax("/api/brand/brand/" + $("#uuid2").val(), "PUT", $('#form-update').serialize(), function (e) {
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
    return myAjax("/api/brand/brand", "GET", {
        "pageNum": pageNum,
        "pageSize": pageSize,
        "name": $("#table_search").val()
    }, function (e) {
        let title = ["name", "btn"]
        tablePage(e, title, "")
    })
}

//删除
function del(o) {
    alter2IsOk("是否确定删除？").then(function (e) {
        if (e.value) {
            return myAjax("/api/brand/brand/" + o, "DELETE", null, function (e) {
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
    return myAjax("/api/brand/brand/" + e, "GET", null, function (e) {
        if (e.success) {
            $("#uuid2").val(e.data.uuid)
            $("#name2").val(e.data.name)
        } else {
            alter2(3, e.msg)
        }
    })
}