package controller

import (
	"encoding/json"
	"fmt"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"github.com/gorilla/websocket"
	"time"
)

//聊天室
type ChatRoom struct {
	Name       string
	OnlineNum  int
	Clients    []*Client
	ClientsMap map[*Client]bool
	MsgChan    chan *models.Msg //聊天室广播通道
}

func (p *ChatRoom) InRoom(conn *Client) {
	msg := &models.Msg{
		KindMsg:    1,
		Msg:        fmt.Sprintf(" -%d- 进入聊天室", conn.User.NO),
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SourceAddr: conn.Conn.RemoteAddr().String(),
		SourceNO:   conn.User.NO,
		User:       conn.User,
	}

	//将客户端的发来的消息推送到 聊天室消息通道
	go func() {
		for {
			select {
			case msg := <-conn.MsgChan:
				p.MsgChan <- msg
			}
		}
	}()

	p.MsgChan <- msg
}

func (p *ChatRoom) OutRoom(conn *Client) {
	msg := &models.Msg{
		KindMsg:    1,
		Msg:        fmt.Sprintf(" -%d- 退出聊天室", conn.User.NO),
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SourceAddr: conn.Conn.RemoteAddr().String(),
		SourceNO:   conn.User.NO,
		User:       conn.User,
	}

	p.MsgChan <- msg
}

//将消息广播给客户端
func (p *ChatRoom) Broadcast(msg *models.Msg) {
	for _, client := range p.Clients {
		//发送给除自己之外的客户端
		if client.Conn.RemoteAddr().String() != msg.SourceAddr {
			SendToClient(client, msg)
		}
	}
}

func (p *ChatRoom) Start() {
	for {
		select {
		case msg := <-p.MsgChan:
			p.Broadcast(msg)
		}
	}
}

func SendToClient(client *Client, msg *models.Msg) {
	b, _ := json.Marshal(msg)
	_ = client.Conn.WriteMessage(websocket.TextMessage, b)
}
