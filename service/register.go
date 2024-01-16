package service

import (
	"fmt"
	"mysingo/api"
	"mysingo/logic"
	"mysingo/model"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	//获取结果
	var user api.RegisterReq
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println("数据绑定失败:", err)
		c.JSON(400, api.RegisterRes{
			Code:    400,
			Message: "数据绑定失败",
			Data:    api.Data{},
		})
		return
	}
	//验证表单
	err = logic.Vail(c, user)
	if err != nil {
		c.JSON(400, api.RegisterRes{
			Code:    400,
			Message: err.Error(),
			Data:    api.Data{},
		})
		return
	}
	//加密密码
	hashpassword, err := logic.HashPassword(user.Password)
	if err != nil {
		c.JSON(400, api.RegisterRes{
			Code:    400,
			Message: err.Error(),
			Data:    api.Data{},
		})
		return
	}
	//创建实例
	modeluser := model.User{
		UserName:       user.UserName,
		Nickname:       user.NickName,
		PasswordDigest: hashpassword,
		Status:         "active",
		Avatar:         "touxiang",
	}
	//存数据库
	err = logic.CreateUser(modeluser)
	if err != nil {
		c.JSON(400, api.RegisterRes{
			Code:    400,
			Message: err.Error(),
			Data:    api.Data{},
		})
		return
	}
	//ok
	c.JSON(200, api.RegisterRes{
		Code:    200,
		Message: "用户创建成功",
		Data:    api.Data{},
	})
}
