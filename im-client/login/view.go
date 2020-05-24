package login

import (
	"fmt"
	"go-im/im-client/game"
)

func TipLogin(pram interface{}) {
	var account string
	fmt.Println("登录提示：请输入账号")
	fmt.Scanln(&account)
	var password string
	fmt.Println("登录提示：请输入密码")
	fmt.Scanln(&password)
	game.Session.Account = account
	loginReq(account, password)
}
