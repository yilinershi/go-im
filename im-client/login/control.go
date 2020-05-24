package login

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
	client.AddRouter(netCommond.TypeLoginResp, &Router{})

	//UI消息
	eventCenter.On(eventCenter.OnTipLogin, TipLogin)
	fmt.Println(" init login modeul success!")
}
