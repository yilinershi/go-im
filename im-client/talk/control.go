package talk

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

	//网络消息
	client.AddRouter(netCommond.TypeFriendTalkSync, &Router{})

	//UI消息
	eventCenter.On(eventCenter.OnGetFriendList, tipInputTalk)    //friend模块获取好友列表，talk模块通知提示输入聊天内容
	fmt.Println("init talk conn success!")
}

