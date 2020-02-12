package db

import (
	"context"
	"github.com/HelloHaiGG/WeChat/common/imongo"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	"log"
	"time"
)

func SyncRecordToMongo(msg *models.Msg) {
	var err error

	cxt, _ := context.WithTimeout(context.Background(), time.Second * 3)

	_, err = imongo.DB.Collection("chat_record").InsertOne(cxt, msg)

	if err != nil {
		log.Fatalf("record insert to mongo err -> %v", err)
	}
}
