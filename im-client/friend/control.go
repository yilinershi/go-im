package friend

import (
	"fmt"
	"go-im/im-common/eventCenter"
	"go-im/im-common/netCommond"
	"go-im/im-common/zinx/iface"
)

var conn iface.IConnection

//初始化客户端的连接
func Init(client iface.IClient) {
	conn = client.GetConn()
	fmt.Println("friend init conn success!")

	//网络消息
	client.AddRouter(netCommond.TypeGetFriendListResp, &Router{})

	//UI消息
	eventCenter.On(eventCenter.OnLoginSuccess, onLoginSuccess)
	eventCenter.On(eventCenter.OnGetFriendList, onGetFriendList)
}

func onLoginSuccess(pram interface{}) {
	fmt.Println("登录成功")
	reqFriendList()
}

func onGetFriendList(pram interface{}) {
	fmt.Println("获取好友列表成功")
	showSelectDia()
}
