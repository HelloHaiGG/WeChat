package router

import (
	"github.com/HelloHaiGG/WeChat/common"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func ChatRouter() *iris.Application {
	app := iris.New()
	app.Use(recover2.New(), logger.New())

	app.OnErrorCode(iris.StatusForbidden, forbidden)
	app.OnErrorCode(iris.StatusInternalServerError, notfound)
	app.OnErrorCode(iris.StatusNotFound, internal)

	app.Get("/", index)

	return app
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
