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

	beego.Router("/querydata.html",&controllers.HomeController{})
}
