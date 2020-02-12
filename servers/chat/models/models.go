package models

import "github.com/HelloHaiGG/WeChat/servers/user/models"

//聊天室
type ChatRoom struct {
	Id       int64  `json:"id" gorm:"-"`
	HNO      int64  `json:"h_no" gorm:"column:H_NO"`
	Number   int    `json:"number" gorm:"number"`
	IsDelete int    `json:"is_delete" gorm:"is_delete"`
	Name     string `json:"name" gorm:"name"`
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
	RId      int64 `json:"r_id" gorm:"r_id"`         //房间列表
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
	Holder     OutHolder   `json:"holder"` //退出聊天室
}

//退出聊天室
type OutHolder struct {
	Out      bool
	RoomName string
}

//创建房间
type CreateChatRoomReq struct {
	Name   string `json:"name"`
	HNO    int64 `json:"hno"` //房主
	Level  int   `json:"level"`
	Number int   `json:"number"`
}

//更新房间
type UpdateChatRoomReq struct {
	HNO    int64 `json:"hno"`
	Number int64 `json:"number"`
}

//通过Name查询房间信息
type QueryRoomByNameReq struct {
	Name string `json:"name"`
}

//返回结果
type QueryRoomByNameRes struct {
	room *ChatRoom
}

//通过ID查询房间信息
type QueryRoomByIdReq struct {
	Id int64 `json:"id"`
}

//返回结果
type QueryRoomByIdRes struct {
	Room *ChatRoom
}

//查询房间内成员的信息
type QueryRoomMembersReq struct {
	RId    int64
	Online int // 1,查询在线 0 不在线
}

//返回结果
type QueryRoomMemberRes struct {
	Users []*models.User
}

//加入到聊天室
type InChatRoomReq struct {
	RoomName string
	NO       int64
	RoomId   int64
}
