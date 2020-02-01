package controller

import (
	"github.com/kataras/iris/mvc"
)

type TestController struct{}

//使用请求方法对函数命名 //GET,POST,DELETE,PUT 等
func (p *TestController) Get() mvc.Result {
	return mvc.Response{
		Code:        200,
		Content:     []byte("ss"),
	}
}
//根据请求方法和URL自动匹配处理方法 例如：GET http://localhost:8080/hello  匹配GetHello
func (p *TestController) GetHello() string {
	return "hello"
}

//使用 BeforeActivation 将请求路径和处理方法一一对应
func (p *TestController)BeforeActivation(b mvc.BeforeActivation)  {
	b.Handle("GET","path/hello","SayHello")
}
func (p *TestController)SayHello() mvc.Result {
	return mvc.Response{
		Code:200,
		Text:"hello world!",
	}
}
