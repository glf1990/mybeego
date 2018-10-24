function TransferString(content)  
{  
    var string = content;  
    try{  
        string=string.replace(/\r\n/g,"")  
        string=string.replace(/\n/g,"");
        string=string.replace(/\r/g,"");  
        string=string.replace(/<br\/>/g,""); 
        string=string.replace(/<br>/g,""); 
    }catch(e) {  
        alert(e.message);  
    }  
    return string;  
} 

$(document).ready(function(){
	//获取数据信息===================
	var env=$("#env").val()
	$.ajax({
                url:"/api/getDatabase",
                type:"post",
                contentType: 'application/json;charset=utf-8',
                data:'{"env":"'+env+'"}',
                timeout:20000,
                dataType:"json",
                success:function (da){
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
    //alert("ceshiceshi")
    var da;
    $("#env").change(function(){
    	$("#database").empty();
		var env=$("#env").val()
		$.ajax({
	                url:"/api/getDatabase",
	                type:"post",
	                contentType: 'application/json;charset=utf-8',
	                data:'{"env":"'+env+'"}',
	                timeout:20000,
	                dataType:"json",
	                success:function (da){
	                	var data=JSON.parse(da)
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
    $("#getdata").click(function(){
    	//$('select').select();
    	var env=$("#env").val()
    	var database=$("#database").val()
    	var sql=$("#sql").val()
    	alert(sql)
    	sql=TransferString(sql)
    	
            $.ajax({
                url:"/api/getdata",
                type:"post",
                contentType: 'application/json;charset=utf-8',
                data:'{"env":"'+env+'","database":"'+database+'","sql":"'+sql+'"}',
                timeout:20000,
                dataType:"json",
                success:function (da) {
                	//转换json对象
                	data=JSON.parse(da)
                	var key = [],value= []
                	//拼接表头
                	k=data[0]
            		for(var p in k){
				        //alert(jsonObj[i].id);  //取json中的值
				        //取它的keyß
				        //key=key+',{"label":'+'"'+p+'"'+',"width":100,"sortable":"default","name":"'+p+'"}'
				        key.push({"label":p,"width":50,"sortable":"default","name":p})
            		}
            		//将数组转换成json字符串
            		key =JSON.stringify(key)
            		var head=JSON.parse(key)
            		
                	for(var i=0;i<data.length;i++){
                		var l=[]
                		for(var p in data[i]){
					        l.push(data[i][p])
            			}
                		value.push(l)
                		//alert(data[i])	
					}
					value =JSON.stringify(value)
					var tbody=JSON.parse(value)
                	//da=data.json
                    //alert("Data Saved: " + data.message+"|"+msg.Success);
                    //da=data.json
					/*                
    				//示例=======================================================================================
				    var oper = [{label:'删除',onclick:function(){
										alert('删除');
								}},{label:'编辑',onclick: function(){
									alert('编辑')	
								}}]
					var tbody = [
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper], 
									["201302","uimaker","小牛","山东济南","山东大学","1989-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper], 
									["201302","uimaker","小牛","山东济南","山东大学","1989-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper], 
									["201302","uimaker","小牛","山东济南","山东大学","1989-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper], 
									["201302","uimaker","小牛","山东济南","山东大学","1989-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper], 
									["201302","uimaker","小牛","山东济南","山东大学","1989-10-18","已审核",oper],
									["201301","admin","熊猫王子","江苏南京","南京林业大学","1982-10-18","已审核",oper]]
					*/		
									
						$('.grid').Grid({
							thead: head,
							tbody: null,
							height:400,
							checkbox: {
								single:true	
							}
							//操作列====================
							/*operator: {
								type : "normal",
								width : 100	
							},
							//排序回调函数
							sortCallBack : function(name,index,type){
								//alert(name+","+index+','+type);
							}*/
							
						});
					
					$('.grid').Grid('addLoading');
					
					/// 模拟异步
					/*setTimeout(function(){
						$('.grid').Grid('setData',tbody, head);
					},2000)*/
				    $('.grid').Grid('setData',tbody, head);
					
	},
    error: function (mm) {
        alert("请求失败，提示:" +mm);
    }
	});
    
    /*//分页=====================================
	$('.pagination').pagination(100,{
		callback: function(page){
			alert(page);	
		},
		display_msg: false
	});
	
	$('.search-box input[type=radio]').click(function(e) {
        if($(this).prop('checked')){
			if($(this).attr('data-define') === 'define'){
				$('.define-input').show();
			}else{
				$('.define-input').hide();
			}
		}
    });*/
        
        });
    
});