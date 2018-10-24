package controllers

import (
	"github.com/astaxie/beego"
)
type Hellotest struct {
	beego.Controller
}

func (this *Hellotest) Get() {
	this.Ctx.WriteString("hello world")
}