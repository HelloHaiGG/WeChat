package controller

import (
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//聊天室管理对象
type RoomManager struct {
	RoomsMap map[string]*ChatRoom
	Rooms    []*ChatRoom
}

//创建聊天室,客户端进入聊天室
func (p *RoomManager) InitRoom(roomName string, client *Client) {
	if _, ok := p.RoomsMap[roomName]; !ok {
		room := &ChatRoom{
			Name:       roomName,
			OnlineNum:  0,
			Clients:    make([]*Client, 0),
			ClientsMap: make(map[*Client]bool),
			MsgChan:    make(chan *models.Msg, 0),
		}
		room.Clients = append(room.Clients, client)
		room.ClientsMap[client] = true
		room.OnlineNum++
		p.RoomsMap[roomName] = room
		p.Rooms = append(p.Rooms, room)
	}

	//开启聊天室广播
	go p.RoomsMap[roomName].Start()

	p.RoomsMap[roomName].InRoom(client)
}

// 客户端进入聊天室
func (p *RoomManager) ClientInRoom(roomName string, client *Client) mvc.Result {
	if room, ok := p.RoomsMap[roomName]; !ok {
		return mvc.Response{Code: iris.StatusInternalServerError,}
	} else {
		room.Clients = append(room.Clients, client)
		room.ClientsMap[client] = true
		room.OnlineNum++
		room.InRoom(client)
	}
	return mvc.Response{
		Code: iris.StatusOK,
	}
}

//客户端离开聊天室
func (p *RoomManager) ClientOutRoom(roomName string, conn *Client) {
	if room, ok := p.RoomsMap[roomName]; ok {
		delete(room.ClientsMap, conn)
		room.OutRoom(conn)
		room.OnlineNum--
	}

	//判断聊天室是否已经没人
	if p.RoomsMap[roomName].OnlineNum <= 0 {
		//删除聊天室
		close(p.RoomsMap[roomName].MsgChan)
		delete(p.RoomsMap, roomName)
	}
}
