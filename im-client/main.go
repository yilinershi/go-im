package main

import (
	"fmt"
	"go-im/im-client/login"
	"go-im/im-client/register"
	"go-im/im-common/eventCenter"
	"go-im/im-common/messageCommond"
	"go-im/im-common/zinx/znet"
	"time"
)

func main() {
	initClient()
	tipInput()
	select {}
}

func initClient() {
	client := znet.NewClient()
	client.Start()

	//登录模块初始化
	login.Init(client.GetConn())
	
	//注册登录路由
	client.AddRouter(messageCommond.TypeLoginResp, &login.Router{})

	register.Init(client.GetConn())
	client.AddRouter(messageCommond.TypeRegisterResp, &register.Router{})
}

func tipInput() {
	//等日志打印完，正常情况下去掉该行
	time.Sleep(time.Second * 1)

	var key int
Loop:
	fmt.Println("\t\t ------欢迎登录多人聊天室--------")
	fmt.Println("\t\t\t 1.登录聊天室")
	fmt.Println("\t\t\t 2.注册用户")
	fmt.Println("\t\t\t 3.退出系统")
	fmt.Println("\t\t\t 请选择1-3")
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("登录聊天")
		eventCenter.Emit(eventCenter.OnTipLogin, nil)
		return
	case 2:
		fmt.Println("注册用户")
		eventCenter.Emit(eventCenter.OnTipRegister, nil)
		return
	case 3:
		fmt.Println("退出聊天")
	default:
		fmt.Println("您输入有误，请重新输入")
		goto Loop
	}
}
