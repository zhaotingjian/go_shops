package controllers

import (
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

//用户登入验证
func (c *CommonController) Prepare() {
}
