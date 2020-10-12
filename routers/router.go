package routers

import (
	"Datarenzheng1010/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//routers:路由
    beego.Router("/", &controllers.MainController{})

	beego.Router("/register",&controllers.RegisterController{})
}
