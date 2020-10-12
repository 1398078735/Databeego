package controllers

import (
	"Datarenzheng1010/db_mysql"
	"Datarenzheng1010/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}


//该方法用于处理用户登录
func (r *RegisterController) Get() {
	//解析用户端请求数据

	//将解析到的数据保存到数据库中

	//将处理结果返回给客户端
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("抱歉，数据错误")
		return
	}
	//r.TplName = "login.html"

	row, err :=db_mysql.AddUser(user)
	if err != nil {
		r.Ctx.WriteString("注册用户信息失败，请重试")
		return
	}
	fmt.Println(row)
	//row代表影响了几行

	if row != -1 {
		r.TplName = "login.html"
	}else {
		r.TplName = "error.html"
	}
}
