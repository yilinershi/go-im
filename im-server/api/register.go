package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/model"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"go-im/im-server/db/user"
)

type RegisterRouter struct {
	znet.BaseRouter
}

func (*RegisterRouter) Handle(request iface.IRequest) {
	msg := request.GetData()
	msgId := request.GetMsgID()
	fmt.Printf("get msg from client, msgId=%d,msg=%s\n", msgId, string(msg))

	req := new(proto.RegisterReq)
	err := json.Unmarshal(msg, req)
	if err != nil {
		fmt.Println("json unmarshal req err, err=", err)
		return
	}

	var resp = func() *proto.RegisterResp {

		if req.Password1 != req.Password2 {
			return &proto.RegisterResp{Error: errorCode.RegisterPasswordError}
		}

		if req.Age <= 14 {
			return &proto.RegisterResp{Error: errorCode.RegisterAgeTooLow}
		}

		dbUser := &model.User{
			Account:  req.Account,
			Password: req.Password1,
			Age:      req.Age,
		}

		putErr := user.AddUser(req.Account, dbUser)
		if putErr != nil {
			return &proto.RegisterResp{Error: errorCode.RegisterDBError}
		}

		return &proto.RegisterResp{Error: errorCode.OK}
	}()

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal 注册 err,err=", err)
		return
	}
	request.GetConnection().SendMsg(netCommond.TypeRegisterResp, jsonData)
	fmt.Printf("注册 resp, msgId=%d,msg=%s\n", netCommond.TypeRegisterResp, string(jsonData))
}
