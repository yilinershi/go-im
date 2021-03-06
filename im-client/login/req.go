package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
)

//发送req
func loginReq(account string, password string) (err error) {

	req := &proto.LoginReq{
		Account:  account,
		Password: password,
	}
	fmt.Printf("login req=%v\n", req)

	jsonData, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		return errors.New("json marshal err")
	}

	//将结构体json byte 发送
	err = conn.SendMsg(netCommond.TypeLoginReq, jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return nil
	}

	return nil
}
