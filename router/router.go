package router

import (
	"github.com/HelloHaiGG/WeChat/common"
	controller2 "github.com/HelloHaiGG/WeChat/servers/chat/controller"
	"github.com/HelloHaiGG/WeChat/servers/user/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/websocket"
)

func ChatRouter() *iris.Application {
	app := iris.New()
	app.Use(recover2.New(), logger.New())

	app.OnErrorCode(iris.StatusForbidden, forbidden)
	app.OnErrorCode(iris.StatusNotFound, notfound)
	app.OnErrorCode(iris.StatusInternalServerError, internal)

	app.Get("/", index)

	app.PartyFunc("/user", func(p router.Party) {
		p.Post("/register", controller.Register).Name = "用户注册"
		p.Post("/login", controller.Login).Name = "用户登录/退出"
	})

	mvc.Configure(app.Party("/friends"), func(a *mvc.Application) {
		a.Handle(new(controller.FriendsController))
	})
	mvc.Configure(app.Party("/ws"), ConfigureMvc)

	return app
}

func ConfigureMvc(a *mvc.Application) {
	ws := websocket.New(websocket.Config{})
	a.Router.Any("/iris-ws.js",websocket.ClientHandler())
	a.Register(ws.Upgrade)
	a.Handle(new(controller2.ChatController))
}

func forbidden(cxt iris.Context) {
	_, _ = cxt.JSON(iris.Map{"code": iris.StatusForbidden, "msg": common.ForbiddenDesc})
}
func notfound(cxt iris.Context) {
	_, _ = cxt.JSON(iris.Map{"code": iris.StatusNotFound, "msg": common.NotFindDesc})
}
func internal(cxt iris.Context) {
	_, _ = cxt.JSON(iris.Map{"code": iris.StatusInternalServerError, "msg": common.InternalDesc})
}
func index(cxt iris.Context) {
	_, _ = cxt.WriteString("Welcome WeChat!")
}
