package controllers

import (
	"Datarenzheng1010/models"
	"github.com/astaxie/beego"
)

type UserKycController struct {
	beego.Controller
}

//浏览器的get请求,用于跳转到实名认证页面
func (u *UserKycController) Get(){
	u.TplName = "user_kyc.html"
}


//form表单的post请求,用于处理实名认证业务
func (u *UserKycController) Post() {
	//1,数据解析,解析前端数据
	var user models.User
	err := u.ParseForm(&user)
	if err != nil {
		u.Ctx.WriteString("数据解析错误")
		return
	}

	//2,把用户的实名认证的信息更新到数据库的用户表当中
	_,err=user.UpdateUser()
	//3,判断实名认证操作结果
	if err != nil{
		u.Ctx.WriteString("实名认证错误")
		return
	}
	//4，跳转或者结果处理
	u.TplName = "home.html"

}