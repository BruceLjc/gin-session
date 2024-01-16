# gin-session for Login/Register
用golang-gin写的session小demo

##session手写实现，没有使用缓存数据库，直接new了个对象，存放在内存和cookie中，有24hour期限

##关于目录
/api	输入输出接口类型，和前端交互的	-类似goframe的/api
/logic	具体细节实现
/model	数据库交互模型&&初始化
/service 路由定义
/session 关于session的细节

##
go mod tidy
go run main.go

很简陋，凑合看吧
关于session的部分网上讲不太清，自己写一个吧，够用了。
