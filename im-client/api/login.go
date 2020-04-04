package api

import (
	"encoding/json"
	"errors"
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
	fmt.Println("get msg from server, msg=",string(msg))
	var loginResp=new(message.LoginResp)
	err:=json.Unmarshal(msg,loginResp)
	if err!=nil{
		fmt.Println("json Unmarshal loginResp err, err=",err)
		return
	}
	if loginResp.Code!=0{
		fmt.Println("login err, err=",loginResp.Error)
		return
	}

	fmt.Println("login success")
}



func LoginReq(userId int,userPassword string) (err error) {
	fmt.Printf("userId=%d,userPassword=%s\n",userId,userPassword)

	var req=&message.LoginReq{
		UserId:userId,
		UserPassword:userPassword,
		UserName:"ming",
	}

	jsonData,err1:=json.Marshal(req)
	if err1!=nil{
		errors.New("json marshal err")
		return
	}

	//将结构体json byte 发送
	err = conn.SendMsg(messageCommond.TypeLoginReq,jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return
	}
	fmt.Println("login req success")
	return
}

