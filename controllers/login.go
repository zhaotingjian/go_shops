package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

// @title 登入页面类型
type LoginController struct {
	CommonController
}

func (c *LoginController) Login() {
	c.TplName = "login.index.tpl"
	if c.Ctx.Request.Method == "POST" {
		//1.获取数据
		username := c.GetString("username")
		password := c.GetString("password")
		//2.过滤数据
		if username == "" {
			c.sendFail(2, "账号不为空！")
		}
		if password == "" {
			c.sendFail(2, "密码不为空！")
		}
		//3.数据判断
		var emp Employees
		o := orm.NewOrm()
		err := o.Raw("SELECT id, user_name, pass_word FROM employees WHERE user_name = ?", username).QueryRow(&emp)
		if err != nil {
			print(err)
			return
		}

		if emp.PassWord != password {
			c.sendFail(2, "密码不正确！")
			return
		}

		//4.放入session
		userInfoStr, err := json.Marshal(emp)
		if err != nil {
			print(err)
			return
		}
		c.SetSession("userinfo", userInfoStr)
		//5.返回数据
		c.sendSuccess(1, "登入成功", "")
		return
	}
}