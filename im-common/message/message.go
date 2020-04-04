package message

type LoginReq struct {
	UserId int
	UserPassword string
	UserName string
}

type LoginResp struct {
	Code int//状态码
	Error string
}

