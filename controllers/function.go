// @title 公共方法
package controllers

import(
	"github.com/astaxie/beego/logs"
)
//输出方式
func print(err error){
	if IsTest {
		panic(err)
	}else{
		logs.Error(err)
	}
}