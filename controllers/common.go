// @title 父类
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// @Description	所有类型都继承公共类型
type CommonController struct {
	beego.Controller
}

// @Description 权限检测
func (c *CommonController) Prepare() {
	//1.判断有没有session

	//2.权限判断
	// do some authenticate stuff
	// do_auth_check
	auth := false
	error := false
	if auth {
		c.Ctx.Abort(403, "Forbidden")
	} else if error {
		c.Ctx.Abort(500, "Internal Server Error")
	}
}

func (c *CommonController) send_success(code int, msg string, data interface{}) {
	type SuccessStruct struct {
		Code int
		Msg  string
		Data interface{}
	}
	mystruct := &SuccessStruct{code, msg, data}
	c.Data["json"] = mystruct
	c.ServeJSON()
}

func (c *CommonController) send_fail(code int, msg string) {
	type FailStruct struct {
		Code int
		Msg  string
	}
	mystruct := FailStruct{Code: code, Msg: msg}
	c.Data["json"] = mystruct
	c.ServeJSON()
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8", 30)
}
