package talk

import (
	"fmt"
	"go-im/im-client/game"
)

func tipInputTalk(pram interface{}) {
	if len(game.Session.Frineds) <= 0 {
		fmt.Println("------您当前暂无好友------")
		return
	}

	fmt.Println("------您的好友列表------")
	var selectFriendAccount string
	for _, v := range game.Session.Frineds {
		fmt.Printf("[%s]\t", v.Account)
	}
	fmt.Println()
	fmt.Println("------输入用户名选择您要对话的人，输入空字符退出聊天----")
	fmt.Scanln(&selectFriendAccount)

	var IsContain = func() bool {
		for _, item := range game.Session.Frineds {
			if item.Account == selectFriendAccount {
				return true
			}
		}
		return false
	}

	//名字输入错误，退出聊天
	if !IsContain() {
		return
	}

	var content string
	fmt.Println("------输入您想对话的内容----")
	fmt.Scanln(&content)
	fmt.Printf("你对好友[%s]说：[%s]\n", selectFriendAccount, content)
	inputTalkMessage(selectFriendAccount, content)
}

func onSyncFriendTalk(from string, content string) {
	fmt.Printf("你的好友[%s]对你说：[%s]\n", from, content)
}
