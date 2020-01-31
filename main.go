package main

import (
	"github.com/HelloHaiGG/WeChat/common/igorm"
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
	igorm.Init("chat_test")
}
