package main

import (
	"github.com/HelloHaiGG/WeChat/common/iredis"
	"github.com/HelloHaiGG/WeChat/config"
	"github.com/HelloHaiGG/WeChat/router"
	"github.com/kataras/iris"
	"log"
)

func main() {
	config.Init()
	iredis.Init(&iredis.IOptions{
		Host:     config.APPCfg.Redis.Host,
		Port:     config.APPCfg.Redis.Port,
		DB:       config.APPCfg.Redis.DB,
		Password: config.APPCfg.Redis.Password,
	})

	//接口

	if err := router.Router().Run(iris.Addr(":2428")); err != nil {
		log.Fatalf("Router Run Err:%v", err)
	}
}
