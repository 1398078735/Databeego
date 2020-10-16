package main

import (
	"Datarenzheng1010/db_mysql"
	_ "Datarenzheng1010/routers"
	"github.com/astaxie/beego"
)

func main(){
	db_mysql.Connect()
	//静态资源文件路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()

}
