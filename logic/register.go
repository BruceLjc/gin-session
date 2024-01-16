package logic

import (
	"errors"
	"fmt"
	"mysingo/api"
	"mysingo/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Vail(c *gin.Context, user api.RegisterReq) (err error) {
	if user.Password != user.Password2 {
		return errors.New("password!=password2")
	}

	//昵称 用户名是否唯一
	var count int64

	// 检查昵称是否唯一
	model.DB.Model(&model.User{}).Where("nickname = ?", user.NickName).Count(&count)
	if count > 0 {
		return errors.New("昵称已存在")
	}

	// 检查用户名是否唯一
	model.DB.Model(&model.User{}).Where("user_name = ?", user.UserName).Count(&count)
	if count > 0 {
		return errors.New("用户名已存在")
	}

	return nil
}

// CreateUser
func CreateUser(user model.User) (err error) {
	result := model.DB.Create(&user)
	if result.Error != nil {
		fmt.Println("创建用户时出错:", result.Error)
		fmt.Println("创建用户时的 SQL 语句:", result.Statement.SQL.String())
		return result.Error
	}
	return nil
}

// HashPassword 将明文密码哈希成密文
func HashPassword(password string) (string, error) {
	// 生成密码哈希，其中 14 表示加密的成本，可以根据需求进行调整
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
