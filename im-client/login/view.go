package login

import (
	"fmt"
	"go-im/im-common/eventCenter"

)

func initEvent()  {
	eventCenter.On(eventCenter.OnTipLogin, TipLogin)
}


func TipLogin(pram interface{}) {
	var account string
	fmt.Println("登录提示：请输入账号")
	fmt.Scanln(&account)
	var password string
	fmt.Println("登录提示：请输入密码")
	fmt.Scanln(&password)
	loginReq(account,password)
}


