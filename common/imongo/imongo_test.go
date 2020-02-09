package imongo

import (
	"context"
	"github.com/HelloHaiGG/WeChat/config"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

type User struct {
	Name string
	Age  int
	Sex  int
	NO   int64
}

func init() {
	config.Init("C:\\Users\\Administrator\\GolandProjects\\WeChat\\config.yaml")
	Init(&IOptions{
		Host:     config.APPCfg.Mongo.Host,
		Port:     config.APPCfg.Mongo.Port,
		DB:       config.APPCfg.Mongo.DB,
		User:     config.APPCfg.Mongo.User,
		Password: config.APPCfg.Mongo.Password,
	})
}

func Test_Mongo(t *testing.T) {
	//集合操作    集合：相当于mysql
	collection := DB.Collection("User")
	u := User{
		Name: "小哥哥",
		Age:  21,
		Sex:  1,
		NO:   456225,
	}
	//插入一条数据
	opt := new(options.InsertOneOptions)
	opt.SetBypassDocumentValidation(true)

	result, err := collection.InsertOne(context.Background(), u,opt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result.InsertedID) //插入数据的_id   可以指定_id,不指定会自动生成_id
}
