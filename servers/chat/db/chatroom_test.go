package db

import (
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/config"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"testing"
)

func init() {
	config.Init("/Users/mac126/workspace/go-project/WeChat/config.yaml")
	igorm.Init("chat")
}

func TestCreateRoom(t *testing.T) {
	if err := CreateRoom(&models.CreateChatRoomReq{
		Name:   "titi",
		HNO:    1580711693,
		Level:  1,
		Number: 1,
	}); err != nil {
		t.Error(err)
	}
}
func TestJoinChatRoom(t *testing.T) {
	if err := JoinChatRoom(&models.InChatRoomReq{
		RoomName: "titi",
		NO:       1580711693,
	}); err != nil {
		t.Error(err)
	}
}
func TestQueryMembers(t *testing.T) {
	if res, err := QueryMembers(&models.QueryRoomMembersReq{
		RId:    1,
		Online: 0,
	}); err != nil {
		t.Error(err)
	} else {
		t.Fatal(res.Users)
	}
}
