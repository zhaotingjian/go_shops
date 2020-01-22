package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//登入页面
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	//首页路由
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/index", &controllers.MainController{}, "*:Index")

}
