package controller

import (
	"encoding/json"
	"fmt"
	"github.com/HelloHaiGG/WeChat/common"
	"github.com/HelloHaiGG/WeChat/common/iredis"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"github.com/HelloHaiGG/WeChat/servers/user/db"
	models2 "github.com/HelloHaiGG/WeChat/servers/user/models"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"net/http"
	"sync"
)

var once sync.Once

type ChatController struct {
	Cxt     iris.Context
	Conn    *websocket.Conn
	Manager *RoomManager
}

var roomManager RoomManager

func (p *ChatController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/room/{room:int64}/{NO:int64}", "InitChatServer").Name = "创建聊天室并加入"
	b.Handle("GET", "/room/in/{room:int64}/{NO:int64}", "InChatRoom").Name = "加入聊天室"
}

//创建聊天室
func (p *ChatController) InitChatServer() mvc.Result {
	once.Do(func() {
		roomManager = RoomManager{
			Rooms: make([]*ChatRoom, 0),
		}
	})

	p.Manager = &roomManager

	//获取聊天室 NO
	roomNO, err := p.Cxt.Params().GetInt64("room")
	//获取成员 NO
	userNO, err := p.Cxt.Params().GetInt64("NO")

	if err != nil {
		return mvc.Response{Code: iris.StatusForbidden,}
	}

	//初始化聊天室名称
	roomName := fmt.Sprintf("room-%d", roomNO)

	//获取成员信息 先在redis中获取,获取不到在mysql中获取
	var user models2.User
	if result, err := iredis.RedisCli.HGet("USER_INFO_KEY", fmt.Sprintf("%d_INFO", userNO)).Result(); err != nil {
		//在mysql中获取
		user, err = db.QueryUserByNumber(userNO)
		if err != nil {
			return mvc.Response{Code: iris.StatusInternalServerError,}
		}
	} else {
		_ = json.Unmarshal([]byte(result), &user)
	}

	//将http请求升级为websocket
	if err := p.upgrade(); err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError,}
	}
	client := &Client{
		RoomName: roomName,
		User:     user,
		Conn:     p.Conn,
		MsgChan:  make(chan *models.Msg, 0),
	}
	//websocket
	go client.ReadMsg()

	//判断改聊天室是否已经存在
	if room, ok := roomManager.RoomMap.Load(roomName); !ok {
		p.Manager.InitRoom(roomName, client)
	} else {
		//判断用户是否已经进入到房间
		if roomManager.UserIsExit(room.(*ChatRoom),client.User.NO){
			return mvc.Response{Code:iris.StatusForbidden,Text:common.UserIsExitsInRoom}
		}
		roomManager.ClientInRoom(roomName, client)
	}

	return mvc.Response{Code: iris.StatusOK,}
}

//加入聊天室
func (p *ChatController) InChatRoom() mvc.Result {

	if err := p.upgrade(); err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}

	//获取聊天室编号
	roomNO, err := p.Cxt.Params().GetInt64("room")
	//获取成员编号
	userNO, err := p.Cxt.Params().GetInt64("NO")

	if err != nil {
		return mvc.Response{Code: iris.StatusForbidden,}
	}

	roomName := fmt.Sprintf("room-%d", roomNO)

	if _, ok := roomManager.RoomMap.Load(roomName); !ok {
		return mvc.Response{Code: iris.StatusForbidden, Text: common.ChatRoomAlreadyExists}
	}

	//获取成员信息
	var user models2.User
	if result, err := iredis.RedisCli.HGet("USER_INFO_KEY", fmt.Sprintf("%d_INFO", userNO)).Result(); err != nil {
		//在mysql中获取
		user, err = db.QueryUserByNumber(userNO)
		if err != nil {
			return mvc.Response{Code: iris.StatusInternalServerError,}
		}
	} else {
		_ = json.Unmarshal([]byte(result), &user)
	}

	client := &Client{
		RoomName: roomName,
		User:     user,
		Conn:     p.Conn,
		MsgChan:  make(chan *models.Msg, 0),
	}
	//websocket
	go client.ReadMsg()

	roomManager.ClientInRoom(roomName, client)

	return mvc.Response{
		Code: iris.StatusOK,
	}
}

//http 升级为 websocket
func (p *ChatController) upgrade() error {
	//将http请求升级为websocket
	upgrade := &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := upgrade.Upgrade(p.Cxt.ResponseWriter(), p.Cxt.Request(), nil)
	if err != nil {
		return err
	}
	p.Conn = conn

	return nil
}
