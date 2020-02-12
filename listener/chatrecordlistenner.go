package listener

import "github.com/HelloHaiGG/WeChat/servers/chat/models"

var RecordChan chan *models.Msg
var RecordBackupChan chan *models.Msg

func init() {
	RecordChan = make(chan *models.Msg, 1000)
}

//将聊天记录保存到mongo
func RecordChanListener() {
	for {
		select {
		case msg := <-RecordChan:

		}
	}
}

//将聊天记录备份到redis 队列
func RecordBackupChanListener() {
	for {
		select {
		case msg := <-RecordBackupChan:

		}
	}
}
