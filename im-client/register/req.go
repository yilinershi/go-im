package register

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
)

//发送req
func registerReq(account string, password1 string, password2 string, age int) (err error) {

	req := &proto.RegisterReq{
		Account:   account,
		Password1: password1,
		Password2: password2,
		Age:       age,
	}
	fmt.Printf("register req=%v\n", req)

	jsonData, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		return errors.New("json marshal err")
	}

	//将结构体json byte 发送
	err = conn.SendMsg(netCommond.TypeRegisterReq, jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return
	}

	return
}
