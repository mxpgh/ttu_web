$("#logout").click(function(event){
    event.preventDefault();
    del_cookie("admin_name");
    window.location.href = "/login/index";
})

$("#basic").click(function(event){
    event.preventDefault();
    window.location.href = "/basic/index";
})

$("#config").click(function(event){
    event.preventDefault();
    window.location.href = "/config/index";
})

$("#status").click(function(event){
    event.preventDefault();
    window.location.href = "/status/index";
})

$("#upload").click(function(event){
    event.preventDefault();
    window.location.href = "/upload/index";
})

function del_cookie(name)
{
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;';
}

$("form[data-type=formAction]").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var action = $(target).attr("action");
    $.post(action, $(target).serialize(), function(ret){
        if(ret.Ret == "0") {
            alert(ret.Reason);
        } else {
            location.href = $(target).attr("form-redirect");
        }
    },"json")
})

$("#uploadform").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var action = $(target).attr("action");
    var data = new FormData($("#uploadform")[0]);

    $.ajax({
        url         : action,
        data        : data,
        cache       : false,
        contentType : false,
        processData : false,
        type        : "POST",
        dataType    : "json",
        enctype     : "multipart/form-data",
        success     :  function(ret){
            if(ret.Ret == "0") {
                alert(ret.Reason);
            } else {
               location.href = $(target).attr("form-redirect");
            }
        }
    });
})