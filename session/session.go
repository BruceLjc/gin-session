package session

import (
	"mysingo/api"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Session struct {
	Session map[interface{}]interface{}
	// 添加过期时间字段
	ExpirationTime time.Time
	UserName       string
	Password       string
}

type Sessions struct {
	Sessions []Session
	Count    int
}

var ISessions = Sessions{
	Sessions: make([]Session, 0),
	Count:    0,
}

// Create
func (s *Sessions) Create(c *gin.Context, user api.LoginReq) (sessionID string, err error) {
	// 生成唯一的 session ID
	sessionID = uuid.New().String()

	// 设置过期时间为24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建 Session 对象并存储
	newSession := Session{
		Session:        make(map[interface{}]interface{}),
		ExpirationTime: expirationTime,
		UserName:       user.UserName,
		Password:       user.Password,
	}

	// 将 Session 存入 Sessions 切片
	s.Sessions = append(s.Sessions, newSession)
	s.Count++

	//将 Session 存入cookie
	c.SetCookie("sessionID", sessionID, 3600, "/", "localhost", false, true)

	return sessionID, nil
}

// Delete
func (s *Sessions) Delete(user api.LoginReq) (err error) {
	// 根据 user 信息删除对应的 session
	for i, session := range s.Sessions {
		if session.Session["username"] == user.UserName && session.Session["password"] == user.Password {
			// 删除 session
			s.Sessions = append(s.Sessions[:i], s.Sessions[i+1:]...)
			s.Count--
			break
		}
	}

	return nil
}

// UpdateTime->24hour
func (s *Sessions) Update(user api.LoginReq) (err error) {
	// 根据 user 信息更新对应 session 的过期时间为24小时
	for i, session := range s.Sessions {
		if session.Session["username"] == user.UserName && session.Session["password"] == user.Password {
			// 更新过期时间为24小时后
			s.Sessions[i].ExpirationTime = time.Now().Add(24 * time.Hour)
			break
		}
	}

	return nil
}

// func (s *Session) SetSessionOnCookie(c *gin.Context, sessionID string) {
// 	// 生成 UUID v4 作为 Session ID
// 	// sessionID := uuid.New().String()

// 	// 将 Session ID 放入 Cookie 中
// 	c.SetCookie("sessionID", sessionID, 3600, "/", "localhost", false, true)
// }
