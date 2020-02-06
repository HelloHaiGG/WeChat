package models

import "github.com/HelloHaiGG/WeChat/servers/user/models"

//聊天室
type ChatRoom struct {
	Id        int64  `json:"id" gorm:"-"`
	NO        int64  `json:"no" gorm:"column:NO"`
	Level     int    `json:"level" gorm:"level"`
	HNO       int64  `json:"h_no" gorm:"column:H_NO"`
	Addr      string `json:"addr" gorm:"addr"`
	Port      string `json:"port" gorm:"port"`
	OnlineNum int    `json:"online_num" gorm:"online_num"`
	IsDelete  int    `json:"is_delete" gorm:"is_delete"`
}

//好友列表
type FriendsList struct {
	Id      int64 `json:"id" gorm:"-"`
	NO      int64 `json:"no" gorm:"column:NO"`
	PNO     int64 `json:"pno" gorm:"column:P_NO"` //好友账号
	POnline int   `json:"p_online" gorm:"p_online"`
}

//房间成员列表
type MembersList struct {
	Id       int64 `json:"id" gorm:"-"`
	NO       int64 `json:"no" gorm:"column:NO"`
	RNO      int64 `json:"no" gorm:"column:R_NO"`    //房间列表
	Identity int   `json:"identity" gorm:"identity"` //成员身份 1.房主 2.成员
}

//消息
type Msg struct {
	KindMsg    int         `json:"kind_msg"` // 1：系统消息 2：用户消息 3：聊天室消息
	Msg        string      `json:"msg"`
	File       string      `json:"file"`
	Time       string      `json:"time"`
	TargetAddr string      `json:"target_addr"`
	TargetNO   int64       `json:"target_no"`
	SourceAddr string      `json:"source_addr"`
	SourceNO   int64       `json:"source_no"`
	User       models.User `json:"user"`
}

//创建房间
type CreateChatRoomReq struct {
	NO    int64 `json:"no"`
	Level int   `json:"level"`
}
