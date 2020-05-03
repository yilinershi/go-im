package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-im/im-common/messageCommond"
	"go-im/im-common/model"
	"go-im/im-common/zinx/iface"
)

var conn iface.IConnection

//初始化客户端的连接
func Init(clientConn iface.IConnection) {
	conn = clientConn
	fmt.Println("client init conn success!")
	initEvent()
}

//发送req
func loginReq(account string,password string) (err error) {

	req:=&model.LoginReq{
		Account: account,
		Password: password,
	}
	fmt.Printf("login req=%v\n", req)

	jsonData, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		errors.New("json marshal err")
		return
	}

	//将结构体json byte 发送
	err = conn.SendMsg(messageCommond.TypeLoginReq, jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return
	}

	return
}
