package main

import (
	"fmt"
	"go-im/im-client/api"
	"go-im/im-common/messageCommond"
	"go-im/im-common/zinx/znet"
	"time"
)

func initClient()  {
	c := znet.NewClient()
	c.Start()
	api.Init(c.GetConn())
	c.AddRouter(messageCommond.TypeLoginResp,&api.LoginRouter{})
}

func main() {
	initClient()

	var key int
Loop:

	fmt.Println("\t\t\t------欢迎登录多人聊天室--------")
	fmt.Println("\t\t\t 1.登录聊天室")
	fmt.Println("\t\t\t 2.注册用户")
	fmt.Println("\t\t\t 3.退出系统")
	fmt.Println("\t\t\t 请选择1-3")
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("登录聊天")
		inputIdAndPwd()
		time.Sleep(time.Second * 100)
		return
	case 2:
		fmt.Println("注册用户")
	case 3:
		fmt.Println("退出聊天")
	default:
		fmt.Println("您输入有误，请重新输入")
		goto Loop
	}
}
//
//func linkServer()(connection *znet.TCPConn){
//	//拿到服务器地址信息
//	hawkServer,err := znet.ResolveTCPAddr("tcp", "localhost:8999")
//	if err != nil {
//
//		os.Exit(1)
//	}
//	//连接服务器
//	connection,err = znet.DialTCP("tcp",nil,hawkServer)
//	if err != nil {
//		fmt.Printf("connect to hawk server error: [%s]",err.Error())
//		os.Exit(1)
//	}
//	return
//}


func inputIdAndPwd() {
	fmt.Println("请输入用户Id")
	var userId int
	fmt.Scanln(&userId)
	fmt.Println("请输入用户密码")
	var userPassword string
	fmt.Scanln(&userPassword)
	api.LoginReq(userId,userPassword)
	time.Sleep(time.Second*200)
}
