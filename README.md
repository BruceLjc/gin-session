# gin-session
用golang-gin写的session小demo

session手写实现，没有使用缓存数据库，直接new了个对象

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

