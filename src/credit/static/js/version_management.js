function TransferString(content)
{
    var string = content;
    try{
        string=string.replace(/\r\n/g,"\n")
        //string=string.replace(/\n/g," ");
        string=string.replace(/\t/g," ");
        string=string.replace(/\r/g,"\n");
        //string=string.replace(/<br\/>/g,"");
        //string=string.replace(/<br>/g,"");
    }catch(e) {
        alert(e.message);
    }
    return string;
}
$(document).ready(function(){
    //add app
    $("#add_app_button").click(function(){
        var app_name=$("#add_app").val()
        var Creator=$("#creator").val()
        $.ajax({
            url:"/api/get_app_info",
            type:"post",
            data:'{"name":"'+app_name+'","creator":"'+Creator+'"}',
            timeout:20000,
            dataType:"json",
            success:function (data){
                alert(data)

               /* var table=document.getElementById("get_app_and_version_info");
                var newRow = table.insertRow(); //创建新行
                var newCell = newRow.insertCell(); //创建新单元格
                newCell.innerHTML = data.name;

                var obj=document.getElementById("belong_product");
                obj.options.add(new Option(data.name,"add"));
                alert(data.name);*/
            },
            error: function (mm) {
                console.dir(mm)
                alert("请求失败，提示:" +mm);
            }
        });
    });

    //add version
    $("#add_version_button").click(function(){
        //var styles = document.getElementById("get_user_id");
        //styles.deleteRow();
        $("#get_user_id").empty();
        var env=$("#env").val()
        var userid=$("#userid").val()
        $.ajax({
            url:"/api/get_user_id_from_platform_id",
            type:"post",
            //contentType: 'application/json;charset=utf-8',
            data:'{"env":"'+env+'","userid":"'+userid+'"}',
            timeout:20000,
            dataType:"json",
            success:function (da){
                var data=JSON.parse(da);
                //$("#userid_msg").html(data[0].user_id);
                //alert(data[0].user_id);
                var table=document.getElementById("get_user_id");
                var newRow = table.insertRow(); //创建新行
                var newCell1 = newRow.insertCell(); //创建新单元格
                var newCell2 = newRow.insertCell();
                newCell1.innerHTML = "platform_user_id";//单元格内的内容
                newCell2.innerHTML = "user_id";
                var newRow = table.insertRow(); //创建新行
                var newCell1 = newRow.insertCell(); //创建新单元格
                var newCell2 = newRow.insertCell();
                newCell1.innerHTML = data[0].platform_user_id;//单元格内的内容
                newCell2.innerHTML = data[0].user_id;
            },
            error: function (mm) {
                console.dir(mm)
                alert("请求失败，提示:" +mm);
            }
        });

    });
});
