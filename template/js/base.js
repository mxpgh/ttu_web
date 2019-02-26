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

$("#container").click(function(event){
    event.preventDefault();
    window.location.href = "/container/index";
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
        xhr: function(){
            myXhr = $.ajaxSettings.xhr();
            if(myXhr.upload){
              myXhr.upload.addEventListener('progress',function(e) {
                if (e.lengthComputable) {
                  var percent = Math.floor(e.loaded/e.total*100);
                  if(percent <= 100) {
                    $("#J_progress_bar").progress('set progress', percent);
                    $("#J_progress_label").html('已上传：'+percent+'%');
                  }
                  if(percent >= 100) {
                    $("#J_progress_label").html('文件上传完毕，请等待...');
                    $("#J_progress_label").addClass('success');
                  }
                }
              }, false);
            }
            return myXhr;
        },
        success     :  function(ret){
            if(ret.Ret == "0") {
                alert(ret.Reason);
            } else {
               location.href = $(target).attr("form-redirect");
            }
        }
    });
})