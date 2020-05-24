package register

import (
	"fmt"
)

func TipRegister(pram interface{}) {
	var account string
	fmt.Println("注册提示：请输入账号")
	fmt.Scanln(&account)

	var password1 string
	fmt.Println("注册提示：请输入密码")
	fmt.Scanln(&password1)

	var password2 string
	fmt.Println("注册提示：请重复密码")
	fmt.Scanln(&password2)

	var age int
	fmt.Println("注册提示：请输入年龄")
	fmt.Scanln(&age)

	registerReq(account, password1, password2, age)
}
