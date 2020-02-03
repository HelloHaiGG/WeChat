package models

import (
	"github.com/HelloHaiGG/WeChat/common/igorm"
	"github.com/HelloHaiGG/WeChat/servers/chat/models"
)

/**
用户注册
*/
func UserRegister(data User) error {
	if err := igorm.DB.Model(User{}).Create(&data).Error; err != nil {
		return err
	}
	return nil
}

/**
用户登录/退出登录
*/

func Login(req LoginReq) error {
	tx := igorm.DB.Begin()

	if err := tx.Model(User{}).Where("NO = ? and is_delete = 0", req.NO).Update(map[string]interface{}{"online": req.IsLogin, "addr": req.Addr}).Error; err != nil {
		tx.Callback()
		return err
	}
	// 更新朋友列表登录状态
	if err := tx.Model(models.FriendsList{}).Where("P_NO = ?", req.NO).Update(map[string]interface{}{"p_online": req.IsLogin}).Error; err != nil {
		tx.Callback()
		return err
	}
	tx.Commit()
	return nil
}

/**
根据用户NO查询用户信息
*/
func QueryUserByNumber(no int64) (User, error) {
	var user User
	if err := igorm.DB.Model(User{}).Where("NO = ? and is_delete = 0", no).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

/**
根据用户昵称查询用户信息
*/
func QueryUserByNickName(nickName string) (User, error) {
	var user User
	if err := igorm.DB.Model(User{}).Where("nick_name = ? and is_delete = 0", nickName).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
