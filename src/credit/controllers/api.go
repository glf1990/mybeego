package controllers

import (
	"common/db_mysql"
	"credit/conf"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type Api_add_risk_control_whitelist_Controller struct {
	beego.Controller
}

func (this *Api_add_risk_control_whitelist_Controller) Post() {
	var m1 map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &m1)//将request body 的值赋值给map
	userid := m1["userid"]
	evn_map := conf.Get_evn_msg()
	ea := evn_map[m1["env"].(string)]
	database := ea.Credit_risk
	strs := strings.Split(strings.Trim(userid.(string), ","), ",")
	db_mysql.Create_db_by_ssh(ea.Database_ip, ea.Database_port, database, ea.Database_userName,
		ea.Database_password, ea.Database_LocalIP, ea.Database_LocalPort, ea.SshServerHost, ea.SshServerPort, ea.SshUserName,
		ea.SshPrivateKeyFilePath, ea.SshKeyPassphrase)
	for i := 0; i < len(strs); i++ {
		str := strings.Split(strings.TrimSpace(strs[i]), " ")
		if len(str) != 2 {
			this.Data["json"] = "data[" + strs[i] + "]Does not meet the requirements ！ "
			this.ServeJSON()
			break
		}
		fmt.Println(str[0], str[1])

		//插入风控白名单
		sql := fmt.Sprintf("insert into risk_control_whitelist_tab(`record_id`,`type`,create_time,update_time,`status`,operator,reason) value('%s',1,'0','0',0,'','');", str[0])

		_, err := db_mysql.Update_getAffectedRowNum(sql)
		if err != nil {
			//str :="{\"msg\":\""+err+"\"}"
			//fmt.Println(err)
			this.Data["json"] = fmt.Sprintf(`{"msg":"error,inser data to risk_control_whitelist_tab error msg:%s,sql:%s"}`, err, sql)
			this.ServeJSON()
		} else {
			sql := fmt.Sprintf("insert into credit_test_score_tab(user_id,`platform_type`,`platform_user_id`,user_name,id_card_no,score,create_time,update_time) value(0,1,'%s','%s','id_card_no',88,'0','0');", str[0], str[1])
			_, err := db_mysql.Update_getAffectedRowNum(sql)
			if err != nil {
				//str :="{\"msg\":\""+err+"\"}"
				//fmt.Println(err)
				this.Data["json"] = fmt.Sprintf(`{"msg":"error,inser data to credit_test_score_tab error msg:%s,sql:%s"}`, err, sql)
				this.ServeJSON()
			}
		}
	}
	this.Data["json"] = fmt.Sprintf(`{"msg":"Success"}`)//server  将输出返回
	this.ServeJSON()

}

type Api_check_credit_user_id_Controller struct {
	beego.Controller
}
func (this *Api_check_credit_user_id_Controller) Post() {
	var m1 map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &m1)//将request body 的值赋值给map ，解析前端传来的jason文件
	platform_user_id := m1["userid"]//
	evn_map := conf.Get_evn_msg()
	ea := evn_map[m1["env"].(string)]
	database := ea.Credit_account
	//strs := strings.Split(strings.Trim(platform_user_id.(string), ","), ",")
	//strs := strings.Trim(platform_user_id.(string))
	strs := platform_user_id.(string)
	db_mysql.Create_db_by_ssh(ea.Database_ip, ea.Database_port, database, ea.Database_userName,
		ea.Database_password, ea.Database_LocalIP, ea.Database_LocalPort, ea.SshServerHost, ea.SshServerPort, ea.SshUserName,
		ea.SshPrivateKeyFilePath, ea.SshKeyPassphrase)
	defer db_mysql.Close_db()
	/*for i := 0; i < len(strs); i++ {
		//str := strings.Split(strings.TrimSpace(strs[i]), " ")
		//puid := platuid[i]
	/*	if len(str) != 1 {
			this.Data["json"] = "data[" + strs[i] + "]Does not meet the requirements ！ "
			this.ServeJSON()
			break
		}*/
//		fmt.Println(strs)
		//puid,err:= strconv.Atoi(str[0])
		//根据platform user id 查看credit userid
		//sql := fmt.Sprintf("select userid from  user_base_info_tab where platform_user_id= %d ", str[0])
		sql := "select user_id from  user_base_info_tab where platform_user_id ="+strs
		fmt.Sprintf(sql)
		pluid,err := db_mysql.Select(sql)//将selcet 得到值赋值给pluid ，类型为row类型的
		if err != nil {
			this.Data["json"] = fmt.Sprintf(`{"msg":"error, check wrong :%s,sql:%s"}`, err, sql)
			this.ServeJSON()
		}else {
			haha,_ := db_mysql.Get_json(pluid)
			//this.Data["json"] = fmt.Sprintf(`{"msg":"%d"}`,  pluid)
			this.Data["json"]=haha// server 返回的jason 格式
			this.ServeJSON()
			}
	//}
	//this.Data["json"] = fmt.Sprintf(`{"msg":"Done!"}`)  //server  将输出返回
	//this.ServeJSON()

}

type Api_get_user_id_Controller struct {
	beego.Controller
}

func (this *Api_get_user_id_Controller) Post() {
	var m1 map[string]interface{}

	json.Unmarshal(this.Ctx.Input.RequestBody, &m1)
	fmt.Println("abc:", m1)
	this.Data["json"] = m1
	this.ServeJSON()

}

type Api_getDatabase_Controller struct {
	beego.Controller
}

func (this *Api_getDatabase_Controller) Post() {
	var m1 map[string]interface{}
	//m1["a"]=this.GetString("username")
	json.Unmarshal(this.Ctx.Input.RequestBody, &m1)
	evn_map := conf.Get_evn_msg()
	ea := evn_map[m1["env"].(string)]

	db_mysql.Create_db_by_ssh(ea.Database_ip, ea.Database_port, "mysql", ea.Database_userName,
		ea.Database_password, ea.Database_LocalIP, ea.Database_LocalPort, ea.SshServerHost, ea.SshServerPort, ea.SshUserName,
		ea.SshPrivateKeyFilePath, ea.SshKeyPassphrase)
	rows, _ := db_mysql.Select("show databases;")
	str, _ := db_mysql.Get_json(rows)
	rows.Close()
	//m2["json"]=str
	//fmt.Println(str)
	this.Data["json"] = str
	this.ServeJSON()

}

type Api_getdataController struct {
	beego.Controller
}

func (this *Api_getdataController) Post() {
	var m1 map[string]interface{}
	//m1["a"]=this.GetString("username")
	json.Unmarshal(this.Ctx.Input.RequestBody, &m1)

	evn_map := conf.Get_evn_msg()

	ea := evn_map[m1["env"].(string)]
	database := m1["database"]

	sql := m1["sql"]
	db_mysql.Create_db_by_ssh(ea.Database_ip, ea.Database_port, database.(string), ea.Database_userName,
		ea.Database_password, ea.Database_LocalIP, ea.Database_LocalPort, ea.SshServerHost, ea.SshServerPort, ea.SshUserName,
		ea.SshPrivateKeyFilePath, ea.SshKeyPassphrase)
	str := strings.ToLower(sql.(string))
	if strings.Contains(str, "insert") || strings.Contains(str, "update") || strings.Contains(str, "delete") {
		affect_num, err := db_mysql.Update_getAffectedRowNum(sql.(string))
		if err != nil {
			//str :="{\"msg\":\""+err+"\"}"
			//fmt.Println(err)
			this.Data["json"] = fmt.Sprintf(`[{"msg":"error,affect num:%d,sql error:%s"}]`, affect_num, err)
			this.ServeJSON()
		} else {

			this.Data["json"] = fmt.Sprintf(`[{"msg":"Success,affect rows num:%d"}]`, affect_num)
			this.ServeJSON()
		}
	} else {
		rows, err := db_mysql.Select(sql.(string))
		if err != nil {
			//str :="{\"msg\":\""+err+"\"}"
			//fmt.Println(err)
			this.Data["json"] = fmt.Sprintf(`[{"sys return msg":"%s"}]`, err)
			this.ServeJSON()
		} else {
			str, _ := db_mysql.Get_json(rows)
			rows.Close()
			//m2["json"]=str
			//fmt.Println(str)
			this.Data["json"] = str
			this.ServeJSON()
		}
	}

}
