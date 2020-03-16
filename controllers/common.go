package controllers

import (
	"fmt"
	"strconv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"regexp"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type CommonController struct {
	beego.Controller
}
const (
	IsTest = true
)

func init() {
	//数据库
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	logs.EnableFuncCallDepth(true)
	logs.Async()
	if IsTest{
		orm.Debug = true
		logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)
	} else {
		orm.Debug = false
		logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	}

}

func (c *CommonController) Prepare() {
	//1.判断登入状态
	url:=[]string{
		"/login",
		"/login/",
		"/login/login",
	}
	check_session:=true
	for _,v:=range url{
		if c.Ctx.Request.URL.Path==v{
			check_session=false
		}
	}
	if check_session {
		v := c.GetSession("userinfo")
		sess:=c.StartSession()
		userinfo:=sess.Get("userinfo")
		if userinfo.id ==0  {
			c.Ctx.Redirect(302, "/login")
		}
		//fmt.Println(v)
		//fmt.Println(json.Unmarshal(v))
	}
	//2.获取菜单权限
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

func (c *CommonController) sendSuccess(code int, msg string, data interface{}) {
	c.Data["json"] = struct {
		Code int
		Msg  string
		Data interface{}
	}{code, msg, data}
	c.ServeJSON()
}

func (c *CommonController) sendFail(code int, msg string) {
	c.Data["json"] = struct {
		Code int
		Msg  string
	}{Code: code, Msg: msg}
	c.ServeJSON()
}
type Page struct{
	HasContents bool
	PageNo      int
	PageSize    int
	TotalPage   int
	TotalCount  int
	FirstPage   bool
	LastPage    bool
	PageUrl     string
	List        interface{}
}
/*1.分页算法；2.url地址的生成*/
func (c *CommonController) pageUtil(col string, where string, pageSize int, list interface{}) Page {
	//1.分页算法
	var err error
	pageNo1:=c.GetString("pageNo")
	pageNo:=1
	if pageNo1 != "" {
		pageNo,err=strconv.Atoi(pageNo1)
		if err != nil {
			print(err)
			return Page{HasContents: false}
		}
	}
	//总数
	var maps []orm.Params
	count := 0
	o := orm.NewOrm()
	_, err = o.Raw("SELECT count(*) as count FROM "+where).Values(&maps)
	if err != nil {
	   print(err)
	   return Page{HasContents: false}
	}
	count,err=strconv.Atoi(maps[0]["count"].(string))
	if err != nil {
	   print(err)
	   return Page{HasContents: false}
	}
	if count <=0{
		return Page{HasContents: false}
	}

	page_start := 0
	page_start = (pageNo - 1) * pageSize
	limit := fmt.Sprintf(" limit %d,%d", page_start, pageSize)
	_, err = o.Raw("SELECT "+col+" FROM "+where+limit).QueryRows(list)
	if err != nil{
	    print(err)
	    return Page{HasContents: false}
	}

	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count / pageSize + 1
	}
	pageUrl:=c.Ctx.Request.URL.Path+"?"+c.Ctx.Request.URL.RawQuery
	re ,err:= regexp.Compile("&pageNo=(.)*")
	if err != nil{
	    print(err)
	    return Page{HasContents: false}
	}
	pageUrl = re.ReplaceAllString(pageUrl, "")
	return Page{HasContents: true, PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp,PageUrl: pageUrl, List: list}
}
//获取请求控制器
func (c *CommonController) getController() string{
	data:=explode(c.Ctx.Request.URL.Path,"/")
	if len(data)>1{
		return data[1]
	}
	return ""
}
//获取请求方法
func (c *CommonController) getAction() string{
	data:=explode(c.Ctx.Request.URL.Path,"/")
	if len(data)>2{
		return data[2]
	}
	return ""
}