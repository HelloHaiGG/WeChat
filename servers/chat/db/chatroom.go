package db

import (
	"errors"
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
	models2 "github.com/HelloHaiGG/WeChat/servers/user/models"
)

//创建房间
func CreateRoom(req *models.CreateChatRoomReq) error {
	if err := igorm.DB.Model(models.ChatRoom{}).Create(&models.ChatRoom{
		HNO:    req.HNO,
		Number: req.Number,
		Name:   req.Name,
	}).Error; err != nil {
		return err
	}
	return nil
}

//更新房间
func UpdateRoom(req *models.UpdateChatRoomReq) error {
	if err := igorm.DB.Model(models.ChatRoom{}).Update(&req).Error; err != nil {
		return err
	}
	return nil
}

//通过房间名称查询房间信息
func QueryRoomByName(req *models.QueryRoomByNameReq) (*models.QueryRoomByIdRes, error) {
	r := new(models.ChatRoom)
	if err := igorm.DB.Model(models.ChatRoom{}).Where("name = ?", req.Name).Take(&r).Error; err != nil {
		return nil, err
	}
	return &models.QueryRoomByIdRes{Room: r}, nil
}

//通过房间Id查询房间信息
func QueryRoomById(req *models.QueryRoomByIdReq) (*models.QueryRoomByIdRes, error) {
	var r *models.ChatRoom
	if err := igorm.DB.Model(models.ChatRoom{}).Where("id = ?", req.Id).Take(&r).Error; err != nil {
		return nil, err
	}
	return &models.QueryRoomByIdRes{Room: r}, nil
}

//查询房间内成员信息
func QueryMembers(req *models.QueryRoomMembersReq) (*models.QueryRoomMemberRes, error) {
	var users []*models2.User
	var ids []*struct {
		NO int64 `json:"no" gorm:"column:NO"`
	}

	//查询房内成员的NO
	if err := igorm.DB.Model(models.MembersList{}).Select("NO").Where("r_id = ?", req.RId).Scan(&ids).Error; err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, errors.New("房间内无成员. ")
	}
	var reqIds []int64
	for _, id := range ids {
		reqIds = append(reqIds, id.NO)
	}

	//查询成员信息
	if err := igorm.DB.Model(models2.User{}).Where("NO in (?) and online = ?", reqIds, req.Online).Scan(&users).Error; err != nil {
		return nil, err
	}

	return &models.QueryRoomMemberRes{Users: users}, nil
}

//成员加入到聊天室
func JoinChatRoom(req *models.InChatRoomReq) error {
	if res, err := QueryRoomByName(&models.QueryRoomByNameReq{Name: req.RoomName}); err != nil {
		return err
	} else {
		if err = igorm.DB.Model(models.MembersList{}).Create(&models.MembersList{NO: req.NO, RId: res.Room.Id}).Error; err != nil {
			return err
		}
	}
	return nil
}
