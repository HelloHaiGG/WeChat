package controller

import (
	"encoding/json"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	models2 "github.com/HelloHaiGG/WeChat/servers/user/models"
	"github.com/gorilla/websocket"
	"time"
)

//客户端
type Client struct {
	RoomName string
	Conn     *websocket.Conn
	User     models2.User
	MsgChan  chan *models.Msg //客户端消息通道
	OutChan  chan *Client     //客户端退出消息监听
}

//func (p *Client) WriteMsg() {
//	//defer func() {
//	//	_ = p.Conn.Close()
//	//}()
//	//for {
//	//	select {
//	//	case msg, ok := <-p.MsgChan:
//	//		if !ok {
//	//			_ = p.Conn.WriteMessage(websocket.CloseMessage, []byte("close."))
//	//			return
//	//		}
//	//		b, _ := json.Marshal(msg)
//	//		_ = p.Conn.WriteMessage(websocket.TextMessage, b)
//	//	}
//	//}
//}

func (p *Client) ReadMsg() {

	var msg []byte
	var err error
	var entityMsg models.Msg

	defer func() {
		_ = p.Conn.Close()
	}()

	for {
		if _, msg, err = p.Conn.ReadMessage(); err != nil {
			//客户端断开链接
			_ = p.Conn.Close()
			p.OutChan <- p
			close(p.MsgChan)
			close(p.OutChan)
			return
		}

		str := string(msg)
		entityMsg.Msg = str
		entityMsg.SourceAddr = p.Conn.RemoteAddr().String()
		entityMsg.KindMsg = 3
		entityMsg.Time = time.Now().Format("2006-01-02 15:04:06")
		entityMsg.SourceNO = p.User.NO
		entityMsg.User = p.User

		_ = json.Unmarshal(msg, entityMsg)

		p.MsgChan <- &entityMsg
	}
}
