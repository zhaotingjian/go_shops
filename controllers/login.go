package controllers

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// @title 登入页面类型
type LoginController struct {
	CommonController
}

func (c *LoginController) Login() {

	if c.Ctx.Request.Method == "POST" {
		//1.获取数据
		username := c.GetString("username")
		password := c.GetString("password")
		//2.过滤数据
		if username == "" {
			c.send_fail(2, "账号不为空！")
			return
		}
		if password == "" {
			c.send_fail(2, "密码不为空！")
			return
		}
		//3.数据判断
		type Employee struct {
			Id       int
			UserName string
			Password string
		}
		var emp Employee

		o := orm.NewOrm()
		err := o.Raw("SELECT id, username, password FROM employee WHERE username = ?", username).QueryRow(&emp)
		if err != nil {
			c.send_fail(1, "该用户不存在！")
			return
		}
		if emp.Password != password {
			c.send_fail(1, "密码不正确！")
			return
		}
		//4.放入session
		v := c.GetSession("userinfo")
		if v == nil {
			c.SetSession("userinfo", &emp)
		}
		fmt.Println(v)
		//5.返回数据
		c.send_success(1, "登入成功", "")
		return
	} else {
		c.TplName = "login.tpl"
	}
}
