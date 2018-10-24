package controllers

import(
	"common/db_mysql"
	"credit/conf"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type Version_control_Controller struct {
	beego.Controller
}
func (c *Version_control_Controller) Get() {
	c.TplName = "version_control.html"
}
type Vc_check_credit_user_id_Controller struct {
	beego.Controller
}
func (this *Vc_check_credit_user_id_Controller) Post() {
	var m1 map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &m1)//将request body 的值赋值给map ，解析前端传来的jason文件
	platform_user_id := m1["userid"]//
	evn_map := conf.Get_evn_msg()
	ea := evn_map[m1["env"].(string)]
	database := ea.Credit_account
	strs := platform_user_id.(string)
	db_mysql.Create_db_by_ssh(ea.Database_ip, ea.Database_port, database, ea.Database_userName,
		ea.Database_password, ea.Database_LocalIP, ea.Database_LocalPort, ea.SshServerHost, ea.SshServerPort, ea.SshUserName,
		ea.SshPrivateKeyFilePath, ea.SshKeyPassphrase)
	defer db_mysql.Close_db()
	sql := "select user_id from  user_base_info_tab where platform_user_id ="+strs
	fmt.Sprintf(sql)
	pluid,err := db_mysql.Select(sql)//将selcet 得到值赋值给pluid ，类型为row类型的
	if err != nil {
		this.Data["json"] = fmt.Sprintf(`{"msg":"error, check wrong :%s,sql:%s"}`, err, sql)
		this.ServeJSON()
	}else {
		haha,_ := db_mysql.Get_json(pluid)
		this.Data["json"]=haha// server 返回的jason 格式
		this.ServeJSON()
	}

}