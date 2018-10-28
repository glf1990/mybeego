package routers

import (
	"credit/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})//主界面
	beego.Router("/navigation", &controllers.NavigationController{})
	beego.Router("/sql_tool", &controllers.Sql_tool_Controller{})//数据库工具1
	beego.Router("/test_tool", &controllers.Test_tool_Controller{})//基本数据查找工具
	beego.Router("/api/getdata", &controllers.Api_getdataController{})
	beego.Router("/api/getDatabase", &controllers.Api_getDatabase_Controller{})
	beego.Router("/api/get_user_id", &controllers.Api_get_user_id_Controller{})
	beego.Router("/api/check_credit_user",&controllers.Api_check_credit_user_id_Controller{})
	beego.Router("/api/add_risk_control_whitelist", &controllers.Api_add_risk_control_whitelist_Controller{})
	//版本控制的url
	beego.Router("/version/version_control", &controllers.Version_control_Controller{})
	beego.Router("/version/check_credit_user",&controllers.Vc_check_credit_user_id_Controller{})
	//zhenbin.huang
	beego.Router("/version_management", &controllers.Version_Management_Controller{})
	beego.Router("/api/get_app_info", &controllers.Api_get_app_info_Controller{})
	beego.Router("/api/get_version_info", &controllers.Api_get_version_info_Controller{})


}
