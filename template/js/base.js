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