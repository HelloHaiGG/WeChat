package controller

import (
	"github.com/HelloHaiGG/WeChat/common"
	"github.com/HelloHaiGG/WeChat/servers/user/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type FriendsController struct {
	Cxt iris.Context
}

func (p *FriendsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/list/{user:int64}/{online:int}", "QueryFriendsList").Name = "获取好友列表"
	b.Handle("POST", "/", "AddFriend").Name = "添加好友"
}
func (p *FriendsController) QueryFriendsList() mvc.Result {
	var req models.QueryFListReq
	var list []*models.User
	var err error
	req.NO, err = p.Cxt.Params().GetInt64("user")
	req.Online, err = p.Cxt.Params().GetInt("online")
	if list, err = models.QueryFList(req); err != nil && err != gorm.ErrRecordNotFound {
		_, _ = p.Cxt.JSON(iris.Map{
			"code": iris.StatusInternalServerError,
			"msg":  common.InternalDesc,
		})
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: "json",
		Object:      list,
	}
}

func (p *FriendsController) AddFriend() mvc.Result {
	var req models.AddFriendReq
	if err := p.Cxt.ReadJSON(&req); err != nil {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}

	if req.NO == 0 || req.PNO == 0 {
		return mvc.Response{
			Code: iris.StatusForbidden,
		}
	}

	//判断 P_NO 是否存在
	if _, err := models.QueryUserByNumber(req.PNO); err != nil && err == gorm.ErrRecordNotFound {
		return mvc.Response{
			Code:    iris.StatusForbidden,
			Content: []byte(common.UserDoesNotExist),
		}
	}

	if err := models.AddFriend(req); err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}
	return mvc.Response{
		Code: iris.StatusOK,
	}
}
