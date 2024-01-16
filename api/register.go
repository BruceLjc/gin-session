package api

import "time"

type RegisterReq struct {
	UserName  string `json:"user_name" binding:"required,min=2,max=30"`
	NickName  string `json:"nick_name" binding:"required,min=1,max=30"`
	Password  string `json:"password" binding:"required,min=6,max=30"`
	Password2 string `json:"password2" binding:"required,min=6,max=30"`
}

type RegisterRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	UserName   string    `json:"user_name" binding:"required,min=2,max=30"`
	NickName   string    `json:"nick_name" binding:"required,min=1,max=30"`
	Password   string    `json:"password" binding:"required,min=6,max=30"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time"`
}
