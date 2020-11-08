package routers

import (
	"Datarenzheng1010/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//routers:路由
    beego.Router("/", &controllers.MainController{})
	//用户注册接口
	beego.Router("/register",&controllers.RegisterController{})
	//用户登录接口
    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/login.html",&controllers.LoginController{})
	//用户上传文件接口
    beego.Router("/home",&controllers.HomeController{})
    beego.Router("/home.html",&controllers.HomeController{})

    beego.Router("/uploadRecord.html",&controllers.HomeController{})
	//查询认证号
	beego.Router("/querydata.html",&controllers.CerDetail{})
    //实名认证
    beego.Router("/user_kyc",&controllers.UserKycController{})
    //短信登录
    beego.Router("/login_sms.html",&controllers.LoginSmsController{})

    //发送验证码登录接口
	beego.Router("/send_sms",&controllers.SendSmsController{})
    //手机号验证码登录功能接口
    beego.Router("/login_sms",&controllers.LoginSmsController{})
}
