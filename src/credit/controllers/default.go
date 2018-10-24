package controllers

import (
	"github.com/astaxie/beego"
)
type Test_tool_Controller struct {
	beego.Controller
}
func (c *Test_tool_Controller) Get() {
	c.TplName = "test_tool.html"
}
type Sql_tool_Controller struct {
	beego.Controller
}
func (c *Sql_tool_Controller) Get() {
	c.TplName = "sql_tool.html"
}


type NavigationController struct {
	beego.Controller
}
func (c *NavigationController) Get() {
	c.TplName = "nav.html"
}
type MainController struct {
	beego.Controller
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "main.html"
}
