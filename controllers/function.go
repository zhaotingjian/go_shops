// @title 公共方法
package controllers

import(
	"strings"
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
//截取字符串
func explode(data string,sep string)[]string{
	 return strings.Split(data, sep)
}
func json_encode(){
	
}
func json_decode(){
	
}