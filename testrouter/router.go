package testrouter

import (
	"fmt"
	"github.com/HelloHaiGG/WeChat/testrouter/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func Router() *iris.Application {
	app := iris.New()

	app.Use(recover2.New())
	app.Use(logger.New())

	//接口形式
	app.Handle("GET","/path",func(cxt iris.Context){
		//形式一
	})
	app.Get("/path/get", func(cxt iris.Context) {
		//形式二
		cxt.JSON(iris.Map{"x": "xx",})
		//cxt.WriteString("sss")
	})

	//路由组 方式一
	say := app.Party("/say")
	//注意：使用 中间件逻辑完成后 context.Next() 继续处理请求
	say.Use(func(context iris.Context) {
		fmt.Println("this middle ware.")
		context.Next()
	})
	say.Get("/hello", func(cxt iris.Context) {
		_, err := cxt.WriteString("Say Hello")
		fmt.Println(err)
	})
	say.Get("/hai", func(cxt iris.Context) {
		cxt.WriteString("Say Hai")
	})
	//路由组 方式二
	app.PartyFunc("/eat", func(p iris.Party) {
		p.Use() //中间件
		p.Get("/banana", func(cxt iris.Context) {
			cxt.WriteString("eat banana.")
		})
		p.Get("/potato", func(cxt iris.Context) {
			cxt.WriteString("eat potato.")
		})
	})

	//mvc 形式  通过 mvc.New() 创建 mvc.Application 对象，application 对象 通过 Handle 绑定 控制器
	mvc.New(app).Handle(new(controller.TestController))
	//mvc 形式  通过 mvc.Configure 设置路由组和配置 控制器
	mvc.Configure(app.Party("/group"), func(a *mvc.Application) {
		a.Handle(new(controller.TestController))
	})
	return app
}
