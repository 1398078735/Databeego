package controllers

import (
	"Datarenzheng1010/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"time"
)

type LoginController struct {
	beego.Controller
}

//直接跳转展示用户登录页面
func (l *LoginController) Get(){
	l.TplName = "login.html"
}

func (l *LoginController) Post() {
	rand.Seed(time.Now().UnixNano())
	time := rand.Intn(1000)
	err := qrcode.WriteFile(" 扫码呀憨憨 "+string(time), qrcode.Medium, 256, "./static/img/qrcode.png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//解析客户端用户提交的登录数据
	var user models.User
	err = l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录信息解析错误")
		return
	}
	//根据解析到的数据执行数据库查询操作
	u,err:=user.QueryUser()
	//判断查询数据库的结果
	if err != nil {
		fmt.Println(err.Error())
		l.Ctx.WriteString("抱歉，用户登录错误")
		return
	}
	//根据查询结果返回客户端相应的信息或者页面跳转
	l.Data["Phone"] = u.Phone//动态数据

	l.TplName = "home.html"
}

