package api

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string
}

type LoginRes struct {
	Code    int
	Message string
}
