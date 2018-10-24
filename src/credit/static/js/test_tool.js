function TransferString(content)  
{  
    var string = content;  
    try{  
        string=string.replace(/\r\n/g,"\n")
        string=string.replace(/\t/g," ");
        string=string.replace(/\r/g,"\n");  
    }catch(e) {
        alert(e.message);  
    }  
    return string;  
} 
$(document).ready(function(){
	$("#addcontrol").click(function(){
		var env=$("#env").val()
		var userids=$("#userid").val()
		userids=TransferString(userids)
		var strs= new Array(); //定义一数组 
		strs=userids.split("\n"); //字符分割 
		var u=""
		for (i=0;i<strs.length ;i++ ) 
		{ 
			if (strs[i] !=""){
				u=u+strs[i]+","
			}
		}
		$.ajax({
	        url:"/api/add_risk_control_whitelist",
	        type:"post",
	        //contentType: 'application/json;charset=utf-8',
	        data:'{"env":"'+env+'","userid":"'+u+'"}',
	        timeout:20000,
	        dataType:"json",
	        success:function (da){
	        	var data=JSON.parse(da)
	        	$("#userid_msg").html(data.msg)
	        },
		    error: function (mm) {
		        console.dir(mm)
		        alert("请求失败，提示:" +mm);
		    }
		});
	});
	$("#getuserid").click(function(){
        var env=$("#env").val()
		var platform_userid=$("#userid").val()
        userids=TransferString(userids)
		var strs= new Array(); //定义一数组 
		strs=userids.split("\n"); //字符分割 
		var u=""
		for (i=0;i<strs.length ;i++ ) 
		{ 
			u=u+strs[i]+","
		} 
		$.ajax({
	        url:"/api/get_user_id",
	        type:"post",
	        //contentType: 'application/json;charset=utf-8',
	        data:'{"userid":"'+u+'"}',
	        timeout:20000,
	        dataType:"json",
	        success:function (da){
	        	alert(da.userid)
	        	return
	        	var data=JSON.parse(da)
	        	$("#database").empty();
	        	for(var i=0;i<data.length;i++){
	        		//alert(data[i].Database)
	        		//$("#database").append(data[i].Database); 
	        		//保证不重复添加
				     
				     //添加子元素
				     $("#database").append("<option>"+data[i].Database+"</option>");
			   	
	        	}
				
				
	        	//转换json对象
	        	//data=JSON.parse(da)
	        },
		    error: function (mm) {
		        console.dir(mm)
		        alert("请求失败，提示:" +mm);
		    }
		});
	});
    $("#checkuserid").click(function(){
        var env=$("#env").val()
       var platform_userid=$("#userid").val()
        $.ajax({
            url:"/api/check_credit_user",
            type:"post",
            data:'{"env":"'+env+'","userid":"'+platform_userid+'"}',
            timeout:20000,
            dataType:"json",
            success:function (da){
            	//一种在前端展示结果的方式
				var data=JSON.parse(da);
                var table=document.getElementById("get_user_id");
                var newRow = table.insertRow(); //创建新行
                //var newCell1 = newRow.insertCell(); //创建新单元格
                var newCell2 = newRow.insertCell();
               // newCell1.innerHTML = "platform_user_id";//单元格内的内容
                newCell2.innerHTML = "user_id";
                var newRow = table.insertRow(); //创建新行
               // var newCell1 = newRow.insertCell(); //创建新单元格
                var newCell2 = newRow.insertCell();
                //newCell1.innerHTML = data[0].platform_user_id;//单元格内的内容
                newCell2.innerHTML = data[0].user_id;
               //$("#userid_msg").html(data.msg)
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
