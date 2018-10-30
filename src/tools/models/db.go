package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
使用orm连接数据库步骤：
//告诉orm使用哪一种数据库
1.注册数据库驱动RegisterDriver(driverName, DriverType)
2.注册数据库RegisterDataBase(aliasName, driverName, dataSource, params ...)
3.注册对象模型RegisterModel(models ...)
4.开启同步RunSyncdb(name string, force bool, verbose bool)
*/

// 在init函数中连接数据库，当导入该包的时候便执行此函数
func init(){
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//dbhost := beego.AppConfig.String("localhost")
	dbhost := "localhost"
	//dbport := beego.AppConfig.String("3306")
	dbport := "3306"
	//dbuser := beego.AppConfig.String("root")
	dbuser :="root"
	//dbpassword := beego.AppConfig.String("glf199077")
	dbpassword :="glf199077"
	//dbname := beego.AppConfig.String("tools")
	dbname :="tools"
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	//orm.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/05blog?charset=utf8")
	orm.RegisterModel(new(User)) // 注册模型，建立User类型对象，注册模型时，需要引入包
	orm.RunSyncdb("default", false, true)
}