package service

import (
	"mysingo/api"
	"mysingo/logic"
	"mysingo/model"

	"github.com/gin-gonic/gin"
)

func LoginHandle(c *gin.Context) {
	var inputuser api.LoginReq
	err := c.ShouldBind(&inputuser)
	if err != nil {
		c.JSON(400, api.LoginRes{
			Code:    400,
			Message: "请输入用户名密码",
		})
		return
	}

	//先根据username从数据找记录
	var modeluser model.User
	var count int64
	model.DB.Where("user_name=?", inputuser.UserName).First(&modeluser).Count(&count)
	if count <= 0 {
		c.JSON(400, api.LoginRes{
			Code:    400,
			Message: "查无此人",
		})
		return
	}
	if !logic.CheckPasswordHash(inputuser.Password, modeluser.PasswordDigest) {
		c.JSON(400, api.LoginRes{
			Code:    400,
			Message: "密码错误",
		})
		return
	}

	//设置sessionid并处理后续
	logic.SetSession(c, inputuser)

	c.JSON(200, gin.H{"message": "login successful !"})

}
