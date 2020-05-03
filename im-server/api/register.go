package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/messageCommond"
	"go-im/im-common/model"
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

	req := new(model.RegisterReq)
	err := json.Unmarshal(msg, req)
	if err != nil {
		fmt.Println("json unmarshal req err, err=", err)
		return
	}

	var resp = func() *model.RegisterResp {

		if req.Password1 != req.Password2 {
			return &model.RegisterResp{Error: errorCode.RegisterPasswordError}
		}

		if req.Age<=14{
			return &model.RegisterResp{Error: errorCode.RegisterAgeTooLow}
		}

		dbUser:=&user.User{
			Account:req.Account,
			Password: req.Password1,
			Age: req.Age,
		}

		putErr := user.PutUser(req.Account, dbUser)
		if putErr != nil {
			return &model.RegisterResp{Error: errorCode.RegisterDBError}
		}

		return &model.RegisterResp{Error: errorCode.OK}
	}()

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal loginResp err,err=", err)
		return
	}
	request.GetConnection().SendMsg(messageCommond.TypeRegisterResp, jsonData)
	fmt.Printf("send msg to client, msgId=%d,msg=%s\n", messageCommond.TypeRegisterResp, string(jsonData))
}
