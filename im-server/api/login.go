package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"go-im/im-server/db/user"
	"go-im/im-server/link"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (*LoginRouter) Handle(request iface.IRequest) {
	msg := request.GetData()
	msgId := request.GetMsgID()
	fmt.Printf("get msg from client, msgId=%d,msg=%s\n", msgId, string(msg))

	req := new(proto.LoginReq)
	err := json.Unmarshal(msg, req)
	if err != nil {
		fmt.Println("json unmarshal loginReq err, err=", err)
		return
	}

	var resp = func() *proto.LoginResp {
		account := req.Account
		user, err := user.GetUserByAccount(account)
		if err != nil {
			return &proto.LoginResp{Error: errorCode.LoginAccountNotExist}
		}

		if req.Password != user.Password {
			return &proto.LoginResp{Error: errorCode.LoginPasswordError}
		}

		u := &link.Client{
			Account: account,
			Conn:    request.GetConnection(),
		}
		link.Clients.Add(u)
		return &proto.LoginResp{Error: errorCode.OK}
	}()

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal loginResp err,err=", err)
		return
	}
	request.GetConnection().SendMsg(netCommond.TypeLoginResp, jsonData)
	fmt.Printf("登录 resp, msgId=%d,msg=%s\n", netCommond.TypeLoginResp, string(jsonData))
}
