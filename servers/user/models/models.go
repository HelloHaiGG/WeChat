package models

//用户表结构
type User struct {
	Id       int64  `json:"id" gorm:"id"`
	NO       int64  `json:"no" gorm:"column:NO"`
	Password string `json:"password" gorm:"password"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	Image    string `json:"image" gorm:"image"`
	Addr     string `json:"addr" gorm:"addr"`
	Port     string `json:"port" gorm:"port"`
	Online   int    `json:"online" gorm:"online"`
	IsDelete int    `json:"is_delete" gorm:"is_delete"`
}
