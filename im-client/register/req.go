package register

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
	
	initEvent()
}

//发送req
func registerReq(account string, password1 string, password2 string,age int) (err error) {

	req := &model.RegisterReq{
		Account:   account,
		Password1: password1,
		Password2: password2,
		Age: age,
	}
	fmt.Printf("register req=%v\n", req)

	jsonData, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		return errors.New("json marshal err")
	}

	//将结构体json byte 发送
	err = conn.SendMsg(messageCommond.TypeRegisterReq, jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return
	}

	return
}
