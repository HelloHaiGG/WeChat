package main

import (
	"github.com/HelloHaiGG/WeChat/common/iredis"
	"github.com/HelloHaiGG/WeChat/config"
)

func main() {
	config.Init()
	iredis.Init(&iredis.IOptions{
		Host:     config.APPCfg.Redis.Host,
		Port:     config.APPCfg.Redis.Port,
		DB:       config.APPCfg.Redis.DB,
		Password: config.APPCfg.Redis.Password,
	})

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


}
