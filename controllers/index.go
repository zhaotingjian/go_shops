package controllers

// @title 首页面类型
type IndexController struct {
	CommonController
}

func (c *IndexController) Index() {
	// menu:=struct{
	// 	name string//名称
	// 	url string//地址
	// 	sort int//排序
	// }{
	// 	"小说管理";
	// 	""
	// }
	c.TplName = "index.index.tpl"
}
