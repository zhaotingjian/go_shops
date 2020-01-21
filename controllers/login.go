package controllers

import (
	"fmt"
	"net/http"
)

type LoginController struct {
	CommonController
}
type JSONStruct struct {
	Code int
	Msg  string
}

func (c *LoginController) Login() {
	re_data := &http.Request{}
	fmt.Println(re_data)
	return
	if re_data.Method == "POST" {
		mystruct := &JSONStruct{0, "hello"}
		c.Data["json"] = mystruct
		c.ServeJSON()
		//接收数据

		//数据库

		//判断

		//返回
	} else {
		c.TplName = "login.tpl"
	}
}
