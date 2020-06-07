package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"go-im/im-server/db/friend"
)

type FriendRouter struct {
	znet.BaseRouter
}

func (*FriendRouter) Handle(request iface.IRequest) {
	msg := request.GetData()
	msgId := request.GetMsgID()
	fmt.Printf("get msg from client, msgId=%d,msg=%s\n", msgId, string(msg))
	
	friendListReq := new(proto.FriendListReq)
	err := json.Unmarshal(msg, friendListReq)
	if err != nil {
		fmt.Println("json unmarshal loginReq err, err=", err)
		return
	}

	var resp = func() *proto.FriendListResp {
		account := friendListReq.Account
		friends, err := friend.GetFriendsByAccount(account)
		if err != nil {
			return &proto.FriendListResp{Error: errorCode.LoginAccountNotExist}
		}
		return &proto.FriendListResp{
			Friends: friends,
			Error:   errorCode.OK,
		}
	}()

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal loginResp err,err=", err)
		return
	}
	request.GetConnection().SendMsg(netCommond.TypeGetFriendListResp, jsonData)
	fmt.Printf("获取好友列表, msgId=%d,msg=%s\n", netCommond.TypeGetFriendListResp, string(jsonData))
}
