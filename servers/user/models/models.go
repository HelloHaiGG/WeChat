package models

//用户表结构
type User struct {
	Id       int64  `json:"id" gorm:"-"`
	NO       int64  `json:"no" gorm:"column:NO"`
	Password string `json:"password" gorm:"password"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	Image    string `json:"image" gorm:"image"`
	Addr     string `json:"addr" gorm:"addr"`
	Port     string `json:"port" gorm:"port"`
	Online   int    `json:"online" gorm:"online"`
	IsDelete int    `json:"is_delete" gorm:"is_delete"`
}

//注册请求
type RegisterReq struct {
	NickName string `json:"nick_name"`
	Password string `json:"password"`
}

//登录
type LoginReq struct {
	IsLogin  int    `json:"is_login"` // 1.登录 0.退出
	NO       int64  `json:"no"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
}

//查询好友列表
type QueryFListReq struct {
	Online int `json:"online"` //在线
	NO     int64 `json:"no"`
}

//添加好友
type AddFriendReq struct {
	NO int64 `json:"no"`
	PNO int64 `json:"p_no"`
}