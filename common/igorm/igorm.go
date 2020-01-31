package igorm

import (
	"fmt"
	"github.com/HelloHaiGG/WeChat/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
)

var DB *gorm.DB

func Init(db string) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		config.APPCfg.Mysql.User,
		config.APPCfg.Mysql.Password,
		config.APPCfg.Mysql.Host,
		strconv.Itoa(config.APPCfg.Mysql.Port),
		db,
	)
	DB, err := gorm.Open("mysql", args)
	if err != nil {
		log.Fatalf("Content DB:%s Error.", db)
	}

	DB.DB().SetMaxIdleConns(10)  //连接池中最大空闲连接数
	DB.DB().SetMaxOpenConns(100) //最大连接数
	DB.SingularTable(true)       //表名单数形式
}
