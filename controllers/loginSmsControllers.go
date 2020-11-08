package controllers

import (
	"Datarenzheng1010/models"
	"github.com/astaxie/beego"
	"time"
)


type LoginSmsController struct {
	beego.Controller
}

//浏览器发起的短信验证码登录请求
func (l *LoginSmsController) Get(){
	l.TplName = "login_sms.html"
}

//短信验证码登录
func (l *LoginSmsController) Post(){
	var smsLogin models.SmsLogin
	err := l.ParseForm(&smsLogin)
	if err != nil {
		l.Ctx.WriteString("验证码登录数据解析失败")
		return
	}

	//先拿手机号,查询user表,看有没有注册
	user := models.User{
		Phone: smsLogin.Phone,
	}
	_ , err = user.QueryUserByphone()
	if err != nil {
		l.Ctx.WriteString("该用户还没有注册")
		return
	}

	//拿用户提交的登录信息到数据库中查询
	sms,err:=models.QuerySmsRecord(smsLogin.BizId,smsLogin.Phone,smsLogin.Code)
	if err != nil {
		l.Ctx.WriteString("验证码登录遇到错误")
		return
	}
	if sms.BizId == "" {//验证码错误,手机号错误
		l.Ctx.WriteString("手机号或者验证码错误")
	}
	now := time.Now().Unix()
	if (now - sms.TimeStamp) > 300000  {
		l.Ctx.WriteString("验证码失效")
	}
	//登录正常,跳转主界面
	l.Data["Phone"] = smsLogin.Phone
	l.TplName = "home.html"


}