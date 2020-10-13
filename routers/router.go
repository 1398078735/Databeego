package routers

import (
	"Datarenzheng1010/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//routers:路由
    beego.Router("/", &controllers.MainController{})

	beego.Router("/register",&controllers.RegisterController{})

    beego.Router("/login",&controllers.LoginController{})
    beego.Router("/login.html",&controllers.LoginController{})

    beego.Router("/home",&controllers.HomeController{})
    beego.Router("/home.html",&controllers.HomeController{})
}
