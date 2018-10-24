package main

import (
	_ "credit/routers"   //下划线的意思是只执行init()函数，未加下划线的表示导入所有的包
	"github.com/astaxie/beego"

)

func main() {
	beego.Run()
}

