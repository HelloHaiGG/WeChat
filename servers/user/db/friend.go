package db

import (
	"bytes"
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	models2 "github.com/HelloHaiGG/WeChat/servers/user/models"
)

func QueryFList(req models2.QueryFListReq) ([]*models2.User, error) {
	var list []*models2.User

	query := new(bytes.Buffer)
	args := make([]interface{}, 0)

	query.WriteString("NO = ?")
	args = append(args, req.NO)

	if req.Online != 0 {
		query.WriteString(" and p_online = ?")
		args = append(args, req.Online)
	}
	if err := igorm.DB.Model(models.FriendsList{}).Where(query.String(), args...).Scan(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func AddFriend(req models2.AddFriendReq) error {
	var friend models.FriendsList
	friend.NO = req.NO
	friend.PNO = req.PNO
	if err := igorm.DB.Model(models.FriendsList{}).Create(&friend).Error; err != nil {
		return err
	}
	return nil
}
