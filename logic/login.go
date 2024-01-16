package logic

import (
	"mysingo/api"
	"mysingo/session"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SetSession(c *gin.Context, user api.LoginReq) {
	//my
	session.ISessions.Create(c, user)
}

// CheckPasswordHash 检查用户输入的密码和数据库中存储的哈希密码是否匹配
func CheckPasswordHash(inputPassword, databasePasswordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(databasePasswordHash), []byte(inputPassword))
	return err == nil
}
