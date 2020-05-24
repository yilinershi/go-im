package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/netCommond"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"go-im/im-common/proto"
	"go-im/im-server/db/user"
)

type LoginRouter struct {
	znet.BaseRouter
}

func (*LoginRouter) Handle(request iface.IRequest) {
	msg := request.GetData()
	msgId := request.GetMsgID()
	fmt.Printf("get msg from client, msgId=%d,msg=%s\n", msgId, string(msg))

	loginReq := new(proto.LoginReq)
	err := json.Unmarshal(msg, loginReq)
	if err != nil {
		fmt.Println("json unmarshal loginReq err, err=", err)
		return
	}

	var resp = func() *proto.LoginResp {
		account := loginReq.Account
		user, err := user.GetUserByAccount(account)
		if err != nil {
			return &proto.LoginResp{Error: errorCode.LoginAccountNotExist}
		}

		if loginReq.Password != user.Password {
			return &proto.LoginResp{Error: errorCode.LoginPasswordError}
		}

		return &proto.LoginResp{Error: errorCode.OK}
	}()

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal loginResp err,err=", err)
		return
	}
	request.GetConnection().SendMsg(netCommond.TypeLoginResp, jsonData)
	fmt.Printf("send msg to client, msgId=%d,msg=%s\n", netCommond.TypeLoginResp, string(jsonData))
}
