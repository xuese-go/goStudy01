$(function () {
    findAll()
})

function findAll() {
    $.ajax("/api/jurisdiction/jurisdictions", {
        type: "GET",
        dataType: 'json',
    }).done(function (e) {
        if (e.success) {
            $(e.data).each(function (i, o) {
                // $("#table-content").append(trs(i, o))
                console.log(o)
            })
        } else {

        }
    })
}