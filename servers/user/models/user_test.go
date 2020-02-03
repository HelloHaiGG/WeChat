package models

import (
	"fmt"
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/config"
	"os"
	"testing"
)

func init() {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	config.Init("C:\\Users\\Administrator\\GolandProjects\\WeChat\\config.yaml")
	igorm.Init("chat")
}
func TestQueryUserByNickName(t *testing.T) {
	user, err := QueryUserByNickName("haha")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}
