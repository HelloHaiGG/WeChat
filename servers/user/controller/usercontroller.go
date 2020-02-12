package controller

import (
	"encoding/json"
	"fmt"
	"github.com/HelloHaiGG/WeChat/common"
	"github.com/HelloHaiGG/WeChat/common/iredis"
	"github.com/HelloHaiGG/WeChat/servers/user/db"
	"github.com/HelloHaiGG/WeChat/servers/user/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"reflect"
	"strings"
)

func Register(cxt iris.Context) {
	var register models.RegisterReq
	if err := cxt.ReadJSON(&register); err != nil {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.ForbiddenDesc,
		})
		return
	}

	if len(register.NickName) < 6 || len(register.NickName) > 18 {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.LengthDoesNotMatch,
		})
		return
	}

	if len(register.Password) < 6 || len(register.Password) > 16 {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.LengthDoesNotMatch,
		})
		return
	}

	//判断用户是否存在
	if user, err := db.QueryUserByNickName(register.NickName); err == nil && !reflect.DeepEqual(user, models.User{}) {
		//存在
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.UserAlreadyExists,
		})
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		//查询出错
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusInternalServerError,
			"msg":  common.InternalDesc,
		})
		return
	}
	//获取用户 NO
	var user models.User
	user.NickName = register.NickName
	user.Password = register.Password
	if NO, err := common.NumberHolder.GetNumber(); err != nil {
		//查询出错
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusInternalServerError,
			"msg":  common.InternalDesc,
		})
		return
	} else {
		user.NO = NO
	}
	user.Port = "8462" //默认端口
	//生成用户数据
	if err := db.UserRegister(user); err != nil {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusInternalServerError,
			"msg":  common.InternalDesc,
		})
		return
	}
	//注册成功
	_, _ = cxt.JSON(iris.Map{
		"code": iris.StatusOK,
		"NO":   user.NO,
	})
}

/**
登录
*/
func Login(cxt iris.Context) {
	var login models.LoginReq
	var err error
	var user models.User
	if err = cxt.ReadJSON(&login); err != nil {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.ForbiddenDesc,
		})
		return
	}

	if login.NO == 0 {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.ForbiddenDesc,
		})
		return
	}

	if login.IsLogin == 1 && len(login.Password) == 0 {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusForbidden,
			"msg":  common.ForbiddenDesc,
		})
		return
	}

	//判断用户是否存在
	if user, err = db.QueryUserByNumber(login.NO); err != nil && err == gorm.ErrRecordNotFound {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusOK,
			"msg":  common.UserDoesNotExist,
		})
		return
	} else if err != nil {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusInternalServerError,
			"msg":  common.InternalDesc,
		})
		return
	}
	//判断密码是否正确
	if strings.Compare(user.Password, login.Password) != 0 {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusOK,
			"msg":  common.ThePasswordIsIncorrect,
		})
		return
	}
	//获取 用户地址
	login.Addr = cxt.RemoteAddr()

	if err := db.Login(login); err != nil {
		_, _ = cxt.JSON(iris.Map{
			"code": iris.StatusInternalServerError,
			"msg":  common.InternalDesc,
		})
		return
	}

	//处理用户信息 redis
	userInfo, _ := json.Marshal(&user)
	if login.IsLogin == 1{
		iredis.RedisCli.HSet("USER_INFO_KEY", fmt.Sprintf("%d_INFO", user.NO), string(userInfo))
	}else{
		iredis.RedisCli.HDel("USER_INFO_KEY",fmt.Sprintf("%d_INFO",user.NO))
	}

	_, _ = cxt.JSON(iris.Map{
		"code": iris.StatusOK,
	})
}
