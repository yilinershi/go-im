package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/message"
	"go-im/im-common/messageCommond"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
)

type LoginRouter struct {
	znet.BaseRouter
}


func (*LoginRouter)Handle(request iface.IRequest){
	msg:=request.GetData()
	msgId:=request.GetMsgID()
	fmt.Printf("get msg from client, msgId=%d,msg=%s\n",msgId,string(msg))

	loginReq:=new(message.LoginReq)
	err:=json.Unmarshal(msg,loginReq)
	if err!=nil{
		fmt.Println("json unmarshal loginReq err, err=",err)
		return
	}

	var resp= func() *message.LoginResp {
		if loginReq.UserId!=100{
			return &message.LoginResp{
				Code:  100,
				Error: "找不到该用用",
			}
		}
		if loginReq.UserPassword!="123456"{
			return &message.LoginResp{
				Code:  200,
				Error: "密码错误",
			}
		}
		return &message.LoginResp{
			Code:  0,
			Error: "success",
		}
	}()

	jsonData,err:=json.Marshal(resp)
	if err!=nil{
		fmt.Println("json marshal loginResp err,err=",err)
		return
	}
	request.GetConnection().SendMsg(messageCommond.TypeLoginResp,jsonData)
	fmt.Printf("send msg to client, msgId=%d,msg=%s\n",messageCommond.TypeLoginResp,string(jsonData))
}