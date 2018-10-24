package conf

type Evn struct {
	SshServerHost     string //跳板机IP或者域名
	SshServerPort     int //跳板机ssh的端口好
	SshUserName       string //登录跳板机的用户名
	SshPrivateKeyFilePath string //私钥的全路径 exported as OpenSSH key from .ppk
	SshKeyPassphrase  string // key file encrytion password //私钥的密码

	Database_LocalIP  string // local localhost ip (client side) 本机的IP地址
	Database_LocalPort  int       // local port used to forward the connection 映射本机的端口
	Database_ip string // remote local ip (server side) 数据库的IP服务器地址
	Database_port int        // remote MySQL port 数据库的端口

	// MySQL access data
	Database_userName string //登录数据库的用户名
	Database_password string //登录数据库的密码
	//dbName string //登录的数据库
	Credit_risk string //风控数据库
	Credit_account string//account server 数据库
}
var evn_map = make(map[string]Evn)

func Get_evn_msg() map[string]Evn  {
	evn_map["test"]=Evn{ "203.117.178.65" ,
		22,
		"ld-linfei_gui",
		"/Users/linfei.gui/.ssh/id_rsa",
		"",

		"127.0.0.1",
		3307 ,
		"10.65.136.147",
		3306  ,
		"root" ,
		"qaws!@1123",
		"credit_risk_control_db" ,
		"credit_account_db",
		  }
	evn_map["uat"]=Evn{ "203.117.178.65" ,
		22, //跳板机ssh的端口好
		"ld-linfei_gui", //登录跳板机的用户名
		"/Users/linfei.gui/.ssh/id_rsa", //私钥的全路径 exported as OpenSSH key from .ppk
		"", // key file encrytion password //私钥的密码

		"127.0.0.1", // local localhost ip (client side) 本机的IP地址
		3308 ,       // local port used to forward the connection 映射本机的端口
		"10.65.136.148", // remote local ip (server side) 数据库的IP服务器地址
		3306  ,      // remote MySQL port 数据库的端口
		// MySQL access data
		"root" ,//登录数据库的用户名
		"qaws!@1123", //登录数据库的密码
		"credit_risk_control_id_db" ,//登录的数据库
		"credit_account_id_db",
	}
	evn_map["staging"]=Evn{ "203.117.178.65",
		22, //跳板机ssh的端口好
		"ld-linfei_gui", //登录跳板机的用户名
		"/Users/linfei.gui/.ssh/id_rsa", //私钥的全路径 exported as OpenSSH key from .ppk
		"", // key file encrytion password //私钥的密码

		"127.0.0.1", // local localhost ip (client side) 本机的IP地址
		3309 ,       // local port used to forward the connection 映射本机的端口
		"10.65.136.149", // remote local ip (server side) 数据库的IP服务器地址
		3306  ,      // remote MySQL port 数据库的端口
		// MySQL access data
		"root" ,//登录数据库的用户名
		"qaws!@1123", //登录数据库的密码
		"credit_risk_control_id_db" ,//登录的数据库
		"credit_account_id_db",
	}
	return evn_map
}


//evn_database_ip = {"test":"10.65.136.147","uat":""}
//evn_database_LocalPort={"test":3306,"uat":3306}
//evn_database_user={"test":"3306","uat":"3306"}
//evn_database_pwd={"test":"qaws!@1123","uat":"3306"}