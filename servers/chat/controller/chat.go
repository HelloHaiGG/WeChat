package controller

import (
	"fmt"
	"github.com/HelloHaiGG/WeChat/common"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"github.com/HelloHaiGG/WeChat/servers/user/db"
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
			RoomsMap: make(map[string]*ChatRoom),
			Rooms:    make([]*ChatRoom, 0),
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

	//判断改聊天室是否已经存在
	if _, ok := roomManager.RoomsMap[roomName]; ok {
		return mvc.Response{Code: iris.StatusForbidden, Text: common.ChatRoomAlreadyExists}
	}

	//获取成员信息
	user, err := db.QueryUserByNumber(userNO)
	if err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError,}
	}

	//将http请求升级为websocket
	if err := p.upgrade(); err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError,}
	}
	client := &Client{
		User:    user,
		Conn:    p.Conn,
		MsgChan: make(chan *models.Msg, 0),
	}
	//websocket
	go client.ReadMsg()

	p.Manager.InitRoom(roomName, client)

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

	if _, ok := roomManager.RoomsMap[roomName]; !ok {
		return mvc.Response{Code: iris.StatusForbidden, Text: common.ChatRoomAlreadyExists}
	}

	//获取成员信息
	user, err := db.QueryUserByNumber(userNO)
	if err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError,}
	}

	client := &Client{
		User:    user,
		Conn:    p.Conn,
		MsgChan: make(chan *models.Msg, 0),
	}
	//websocket
	go client.ReadMsg()

	roomManager.ClientInRoom(roomName, client)

	//p.Manager.InitRoom(":6666", p.Conn)
	return mvc.Response{
		Code: iris.StatusOK,
	}
}

//退出聊天室
func (p *ChatController) OutChatRoom() mvc.Result {

	if err := p.upgrade(); err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}

	//p.Manager.ClientOutRoom(":6666", &p.Conn)

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
