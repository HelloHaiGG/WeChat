package models

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
