package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//首页路由
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
}
