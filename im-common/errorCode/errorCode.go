package errorCode


type ErrorBase struct {
	Code int
	Desc string
}

var (
	OK= &ErrorBase{Code: 0, Desc: "成功"}
	LoginAccountNotExist  = &ErrorBase{Code: 1, Desc: "登录用户不存在"}
	LoginPasswordError = &ErrorBase{Code: 2, Desc: "登录密码错误"}
	RegisterAccountExist = &ErrorBase{Code: 3, Desc: "注册用户已存在"}
	RegisterPasswordError = &ErrorBase{Code: 4, Desc: "注册输入两次密码不同"}
	RegisterDBError=&ErrorBase{Code: 5, Desc: "注册写入数据库出错"}
	RegisterAgeTooLow=&ErrorBase{Code: 6, Desc: "注册用户太小，不许玩游戏"}
)
