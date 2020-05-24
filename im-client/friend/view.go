package friend

import (
	"fmt"
	"go-im/im-client/game"
)

func showSelectDia() {
	fmt.Println("------您的好友列表------")
	var selectFriendAccount string
	for _, v := range game.Session.Frineds {
		fmt.Printf("[%s]\t", v.Account)
	}
	fmt.Println()
	fmt.Println("------输入用户名选择您要对话的人----")
	fmt.Scanln(&selectFriendAccount)

	var dia string
	fmt.Println("------输入您想对话的内容----")
	fmt.Scanln(&dia)
	fmt.Printf("你对好友[%s]说：[%s]\n", selectFriendAccount, dia)
}
