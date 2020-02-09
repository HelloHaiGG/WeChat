package main

import (
	"fmt"
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/common/imongo"
	"github.com/HelloHaiGG/WeChat/common/iredis"
	"github.com/HelloHaiGG/WeChat/config"
	"github.com/HelloHaiGG/WeChat/router"
	"github.com/kataras/iris"
)

func main() {
	config.Init()
	iredis.Init(&iredis.IOptions{
		Host:     config.APPCfg.Redis.Host,
		Port:     config.APPCfg.Redis.Port,
		DB:       config.APPCfg.Redis.DB,
		Password: config.APPCfg.Redis.Password,
	})
	imongo.Init(&imongo.IOptions{
		Host:       config.APPCfg.Mongo.Host,
		Port:       config.APPCfg.Mongo.Port,
		DB:         config.APPCfg.Mongo.DB,
		User:       config.APPCfg.Mongo.User,
		Password:   config.APPCfg.Mongo.Password,
		AuthSource: config.APPCfg.Mongo.AuthSource,
		TimeOut:    config.APPCfg.Mongo.Timeout,
	})

	igorm.Init("chat")
	////启动路由
	//if err := testrouter.Router().Run(iris.Addr(":2428")); err != nil {
	//	log.Fatalf("Router Run Err:%v", err)
	//}

	//启动路由 二
	//listener, err := net.Listen("tcp", ":2428")
	//if err != nil {
	//	log.Fatalf("Router Run Err:%v", err)
	//}
	//err = testrouter.Router().Run(iris.Listener(listener))
	//if err != nil {
	//	log.Fatalf("Router Run Err:%v", err)
	//}

	if err := router.ChatRouter().Run(iris.Addr(":2428")); err != nil {
		fmt.Println(err)
	}
}
