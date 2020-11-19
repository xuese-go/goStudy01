$(function () {
    //分页
    page()
    //新增
    $("#user-btn-save").on("click", function () {
        $.ajax("/api/user/user", {
            type: "POST",
            dataType: 'json',
            data: $('#user-form-save').serialize()
        }).done(function (e) {
            if (e.success) {
                alter2(1, "成功")
                $("#user-btn-to-save").click()
                $('#user-form-save')[0].reset()
                page()
            } else {
                alter2(3, e.msg)
            }
        }).fail(function (err) {

        })
    })
})

//分页查询
function page() {
    $.ajax("/api/user/users", {
        type: "GET",
        dataType: 'json',
        data: {
            "pageNum": 1,
            "pageSize": 10,
            "account": $("#table_search").val()
        }
    }).done(function (e) {
        console.log(e)
        if (e.success) {

        } else {

        }
    }).fail(function (err) {

    })
}

//tr模板
function trs(e) {
    return '<tr>'
        + '    <td>1</td>'
        + '    <td>lkfjldjpoawhpfmhpliadsfljoo</td>'
        + '    <td><span class="badge bg-danger">管理员</span><span class="badge bg-primary">普通</span></td>'
        + '    <td><span class="badge bg-success">正常</span><span class="badge bg-secondary">停用</span></td>'
        + '    <td>'
        + '        <button type="button" class="btn btn-danger btn-xs">删除</button>'
        + '        <button type="button" class="btn btn-warning btn-xs">修改/查看</button>'
        + '    </td>'
        + '</tr>'
}