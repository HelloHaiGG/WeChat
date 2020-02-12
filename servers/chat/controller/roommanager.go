package controller

import (
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"sync"
)

//聊天室管理对象
type RoomManager struct {
	//RoomsMap map[string]*ChatRoom
	RoomMap sync.Map
	Rooms   []*ChatRoom
}

//创建聊天室,客户端进入聊天室
func (p *RoomManager) InitRoom(roomName string, client *Client) {
	if _, ok := p.RoomMap.Load(roomName); !ok {
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
		p.RoomMap.Store(roomName, room)
		p.Rooms = append(p.Rooms, room)
	}

	//开启聊天室广播
	go func() {
		room, _ := p.RoomMap.Load(roomName)
		room.(*ChatRoom).Start()
	}()

	room, _ := p.RoomMap.Load(roomName)

	//发放消息
	room.(*ChatRoom).InRoom(client)
}

// 客户端进入聊天室
func (p *RoomManager) ClientInRoom(roomName string, client *Client) mvc.Result {
	if room, ok := p.RoomMap.Load(roomName); !ok {
		return mvc.Response{Code: iris.StatusInternalServerError,}
	} else {
		r := room.(*ChatRoom)

		r.Clients = append(r.Clients, client)
		r.ClientsMap[client] = true
		r.OnlineNum++
		r.InRoom(client)
	}
	return mvc.Response{
		Code: iris.StatusOK,
	}
}

//客户端离开聊天室
func (p *RoomManager) ClientOutRoom(roomName string, addr string) {
	//取到对应的房间
	if room, ok := p.RoomMap.Load(roomName); ok {
		r := room.(*ChatRoom)
		//再找到对应的连接
		for k, _ := range r.ClientsMap {
			if k.Conn.RemoteAddr().String() == addr {
				delete(r.ClientsMap, k)
			}
		}
		r.OnlineNum--
	}

	//判断聊天室是否已经没人
	if room, ok := p.RoomMap.Load(roomName); ok {
		r := room.(*ChatRoom)
		if r.OnlineNum <= 0 {
			close(r.MsgChan)
			p.RoomMap.Delete(roomName)
		}
	}
}

//判断用户是否已经存在于房间
func (p *RoomManager) UserIsExit(room *ChatRoom, no int64) bool {

	if room.OnlineNum <= 0 {
		return false
	} else {
		for _, client := range room.Clients {
			if client.User.NO == no {
				return true
			}
		}
		return false
	}
}
