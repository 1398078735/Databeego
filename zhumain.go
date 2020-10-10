package main

import (
	"Datarenzheng1010/db_mysql"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	db_mysql.Connect()
	//静态资源文件路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}
