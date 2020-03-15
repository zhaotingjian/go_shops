package routers

import (
	"test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//登入页面
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	//主页路由
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/main", &controllers.MainController{}, "*:Index")
	//首页
	beego.Router("/index", &controllers.IndexController{}, "*:Index")
	//小说模块
	beego.Router("/novel/index", &controllers.NovelController{}, "*:Index")
	beego.Router("/novel/", &controllers.NovelController{}, "*:Index")
	beego.Router("/novel", &controllers.NovelController{}, "*:Index")
	beego.Router("/novel/add", &controllers.NovelController{}, "*:Add")
	beego.Router("/novel/edit", &controllers.NovelController{}, "*:Edit")
	beego.Router("/novel/del", &controllers.NovelController{}, "*:Del")
	//员工模块
	beego.Router("/employee", &controllers.EmployeeController{}, "*:Index")
	beego.Router("/employee/", &controllers.EmployeeController{}, "*:Index")
	beego.Router("/employee/index", &controllers.EmployeeController{}, "*:Index")
	beego.Router("/employee/add", &controllers.EmployeeController{}, "*:Add")
	beego.Router("/employee/edit", &controllers.EmployeeController{}, "*:Edit")
	beego.Router("/employee/del", &controllers.EmployeeController{}, "*:Del")

}
