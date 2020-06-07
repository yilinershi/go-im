package main

import (
	"fmt"
	"go-im/im-common/netCommond"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"go-im/im-server/api"
	"go-im/im-server/db"
)

func main() {
	//创建服务器句柄
	s := znet.NewServer()
	db.Init()

	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	//注册路由
	s.AddRouter(netCommond.TypeLoginReq, &api.LoginRouter{})
	s.AddRouter(netCommond.TypeRegisterReq, &api.RegisterRouter{})
	s.AddRouter(netCommond.TypeGetFriendListReq, &api.FriendRouter{})
	s.AddRouter(netCommond.TypeFriendTalkInput, &api.TalkRouter{})
	//启动服务
	s.Serve()

}

//当客户端断开连接的时候的hook函数
func OnConnectionLost(conn iface.IConnection) {
	fmt.Println("OnConnectionLost")
}

//当客户端建立连接的时候的hook函数
func OnConnectionAdd(conn iface.IConnection) {
	fmt.Println("OnConnectionAdd")

}
