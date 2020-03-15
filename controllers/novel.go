package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// @title 小说模块
type NovelController struct {
	CommonController
}
type Novel struct {
	Id         int
	Title      string
	Author     string
	Des        string
	Status     int
	CreateTime string
	EditTime   string
	IsDel      int
}

func init() {
	logs.Debug("novel.init")
}

//列表
func (c *NovelController) Index() {
	//搜索框
	var novel []Novel
	col := "id, title,author,des,status,"
	col += "if(create_time>0,FROM_UNIXTIME(create_time,'%Y-%m-%d %H:%i:%s'),'') as create_time,"
	col += "if(edit_time>0,FROM_UNIXTIME(edit_time,'%Y-%m-%d %H:%i:%s'),'') as edit_time"
	where := "novel WHERE is_del = 1"

	Search := make(map[string]string)
	Search["id"] = c.GetString("id")
	Search["title"] = c.GetString("title")
	Search["author"] = c.GetString("author")
	Search["des"] = c.GetString("des")
	Search["status"] = c.GetString("status")
	Search["create_time"] = c.GetString("create_time")
	Search["edit_time"] = c.GetString("edit_time")
	for k, v := range Search {
		if v != "" {
			where += fmt.Sprintf(" and %v='%v'", k, v)
		}
	}

	c.Data["Search"] = Search
	Page := c.pageUtil(col, where, 10, &novel)
	c.Data["Page"] = Page
	c.TplName = "novel.index.tpl"
}

//增加
func (c *NovelController) Add() {
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
	c.TplName = "novel.add.tpl"
}
func (c *NovelController) Edit() {
	// menu:=struct{
	// 	name string//名称
	// 	url string//地址
	// 	sort int//排序
	// }{
	// 	"小说管理";
	// 	""
	// }
	c.TplName = "novel.edit.tpl"
}
func (c *NovelController) Del() {
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
