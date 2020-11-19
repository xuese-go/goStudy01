$(function () {
    let t = window.localStorage.getItem("xueSeToken")
    if (!t) {
        window.location.href = window.origin + "/"
    }
//    全局设置
    $.ajaxSetup({
        header: {xueSeToken: t},
        timeout: 2000,
        dataType: "json"
    });
})