package main

import (
	_ "4credit/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"4credit/models"
)


func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}


func main() {
	beego.Run()
}

