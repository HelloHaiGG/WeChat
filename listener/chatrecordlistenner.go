package listener

import (
	"github.com/HelloHaiGG/WeChat/servers/chat/db"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
)

var RecordChan chan *models.Msg
//var RecordBackupChan chan *models.Msg

func init() {
	RecordChan = make(chan *models.Msg, 1000)
}

//将聊天记录保存到mongo
func RecordChanListener() {
	for {
		select {
		case msg, ok := <-RecordChan:
			if ok {
				db.SyncRecordToMongo(msg)
			} else {
				close(RecordChan)
			}
		}
	}
}

//将聊天记录备份到redis 队列
//思路：判断聊天室内未上线的玩家，将聊天室信息推送的队列。玩家上线后，将信息推送给玩家，并删除队列
//func RecordBackupChanListener(msg *models.Msg) {
//	for {
//		select {
//		case msg,ok := <-RecordBackupChan:
//			if ok{
//				//
//
//			}else{
//				close(RecordBackupChan)
//			}
//		}
//	}
//}
