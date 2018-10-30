package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
/*
使用orm连接数据库步骤：
//告诉orm使用哪一种数据库
1.注册数据库驱动RegisterDriver(driverName, DriverType)
2.注册数据库RegisterDataBase(aliasName, driverName, dataSource, params ...)
3.注册对象模型RegisterModel(models ...)
4.开启同步RunSyncdb(name string, force bool, verbose bool)
*/
func Init() {//初始化数据库连接，并将用户注册到orm里去
	dbhost := beego.AppConfig.String("mydbhost")
	dbport := beego.AppConfig.String("mydbport")
	dbuser := beego.AppConfig.String("mydbuser")
	dbpassword := beego.AppConfig.String("mydbpassword")
	dbname := beego.AppConfig.String("mydbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User))
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}

type Config struct {
	Id    int
	Name  string
	Value string
}

func (m *Config) TableName() string {
	return TableName("config")
}