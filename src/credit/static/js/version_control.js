
$(document).ready(function(){
    $("#checkuserid2").click(function(){
        var env=$("#env").val()
       var platform_userid=$("#userid").val()
        $.ajax({
            url:"/version/check_credit_user",
            type:"post",
            data:'{"env":"'+env+'","userid":"'+platform_userid+'"}',
            timeout:20000,
            dataType:"json",
            success:function (da){
                var data2=JSON.parse(da);
                //使用layui在前端的展示数据表格的方式
                layui.use('table', function(){
                    var table = layui.table;

                   //第一个实例
                    table.render({
                        elem: '#demo'//html 文件使用的id
                        ,height: 312
                        ,page: true //开启分页
                        ,cols: [[ //表头 ,field 为 传过来的json 文件
                            {field: 'user_id', title: 'userid', width:80, sort: true, fixed: 'left'}
                        ]]
                        ,data:data2
                    });


                });
            	//一种在前端展示结果的方式
/*				var data=JSON.parse(da);
                var table=document.getElementById("get_user_id");
                var table1=form.on

                var newRow = table.insertRow(); //创建新行
                var newCell2 = newRow.insertCell();
                newCell2.innerHTML = "user_id";
                var newRow = table.insertRow(); //创建新行
                var newCell2 = newRow.insertCell();
                newCell2.innerHTML = data[0].user_id;*/

            },
            error: function (mm) {
                console.dir(mm)
                alert("请求失败，提示122222:" +mm);
            }
        });
    });
	$("#now").click(function(){
		//var time=$("#timestamp").val()
		var timestamp = Date.parse(new Date());
		timestamp = timestamp / 1000;
		$("#time_msg").html(timestamp)
	});
	$("#time_to").click(function(){
		var time=$("#timestamp").val()
		var timestamp = Date.parse(new Date(time));
		timestamp = timestamp / 1000;
		$("#time_msg").html(timestamp)
	});
	$("#timestamp_to").click(function(){
		Date.prototype.format = function(format) {
        var date = {
           "M+": this.getMonth() + 1,
           "d+": this.getDate(),
           "h+": this.getHours(),
           "m+": this.getMinutes(),
           "s+": this.getSeconds(),
           "q+": Math.floor((this.getMonth() + 3) / 3),
           "S+": this.getMilliseconds()
        };
        if (/(y+)/i.test(format)) {
               format = format.replace(RegExp.$1, (this.getFullYear() + '').substr(4 - RegExp.$1.length));
        }
        for (var k in date) {
               if (new RegExp("(" + k + ")").test(format)) {
                      format = format.replace(RegExp.$1, RegExp.$1.length == 1? date[k] : ("00" + date[k]).substr(("" + date[k]).length));
               }
        }
        return format;
 		}
		var time=$("#timestamp").val()
		var newDate = new Date();
		newDate.setTime(time * 1000);
		$("#time_msg").html(newDate.format('yyyy-MM-dd h:m:s'))
	});
});
