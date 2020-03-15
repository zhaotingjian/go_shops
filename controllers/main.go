package controllers

// @title 首页面类型
type MainController struct {
	CommonController
}

func (c *MainController) Index() {
	c.TplName = "main.index.tpl"
}
