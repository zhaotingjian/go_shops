package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// @title 小说模块
type EmployeeController struct {
	CommonController
}
type Employees struct {
	Id         int
	UserName   string
	PassWord   string
	TrueName   string
	NickTime   string
	Sex        int
	Age        int
	Mobile     int64
	Card       int64
	CreateTime string
	EditTime   string
	Status     int
	IsTest     int
}

func init() {
	//注册数据
	orm.RegisterModel(new(Employees))
}

//列表
func (c *EmployeeController) Index() {
	//搜索框
	var emp []Employees
	col := "id, user_name, true_name, sex, mobile, card, create_time, age, login_time, nick_name, edit_time, status, is_test"
	where := "employees where 1=1"

	Search := make(map[string]string)
	Search["id"] = c.GetString("id")
	Search["user_name"] = c.GetString("user_name")
	Search["true_name"] = c.GetString("true_name")
	Search["sex"] = c.GetString("sex")
	Search["mobile"] = c.GetString("mobile")
	Search["card"] = c.GetString("card")
	Search["create_time"] = c.GetString("create_time")
	Search["age"] = c.GetString("age")
	Search["login_time"] = c.GetString("login_time")
	Search["nick_name"] = c.GetString("nick_name")
	Search["edit_time"] = c.GetString("edit_time")
	Search["status"] = c.GetString("status")
	Search["is_test"] = c.GetString("is_test")
	for k, v := range Search {
		if v != "" {
			where += fmt.Sprintf(" and %v='%v'", k, v)
		}
	}
	//当前页
	Page := c.pageUtil(col, where, 10, &emp)

	c.Data["Search"] = Search
	c.Data["Page"] = Page

	c.TplName = "employee.index.tpl"
}

//增加
func (c *EmployeeController) Add() {
	if c.Ctx.Request.Method == "POST" {
		orm.Debug = true
		var novel Novel
		//1.获取数据
		time := time.Now().Unix()
		status_str := c.GetString("status")
		if status_str == "" {
			c.sendFail(2, "是否显示为必选项！")
			return
		}
		status, err := strconv.Atoi(status_str)
		if err != nil {
			panic(err)
		}
		novel.Title = c.GetString("title")
		novel.Author = c.GetString("author")
		novel.Des = c.GetString("desc")
		novel.Status = status
		novel.CreateTime = fmt.Sprintf("%d", time)
		novel.EditTime = fmt.Sprintf("%d", time)
		novel.IsDel = 1
		//2.过滤数据
		if novel.Title == "" {
			c.sendFail(2, "标题不为空！")
			return
		}
		if novel.Author == "" {
			c.sendFail(2, "作者不为空！")
			return
		}
		//3.数据判断
		o := orm.NewOrm()
		id, err := o.Insert(&novel)
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
		c.sendSuccess(1, "操作成功！", "")
		return
	}
	c.TplName = "employee.add.tpl"
}
func (c *EmployeeController) Edit() {
	// menu:=struct{
	// 	name string//名称
	// 	url string//地址
	// 	sort int//排序
	// }{
	// 	"小说管理";
	// 	""
	// }
	c.TplName = "employee.edit.tpl"
}
func (c *EmployeeController) Del() {
	// menu:=struct{
	// 	name string//名称
	// 	url string//地址
	// 	sort int//排序
	// }{
	// 	"小说管理";
	// 	""
	// }
	//c.TplName = "novel.edit.tpl"
}
