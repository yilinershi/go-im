package friend

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-im/im-client/game"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
)

//发送req
func reqFriendList() (err error) {
	req := &proto.FriendListReq{
		Account: game.Session.Account,
	}
	fmt.Printf("getFriendList req=%v\n", req)

	jsonData, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		return errors.New("json marshal err")
	}

	//将结构体json byte 发送
	err = conn.SendMsg(netCommond.TypeGetFriendListReq, jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return nil
	}
	fmt.Println("client send msg success: ")
	return nil
}
