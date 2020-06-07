package talk

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-im/im-client/game"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
)

//发送req
func inputTalkMessage(to string, content string) (err error) {

	req := &proto.TalkMessageInput{
		From:  game.Session.Account,
		To:to,
		Content: content,
	}
	fmt.Printf("inputTalkMessage req=%v\n", req)

	jsonData, jsonErr := json.Marshal(req)
	if jsonErr != nil {
		return errors.New("json marshal err")
	}

	//将结构体json byte 发送
	err = conn.SendMsg(netCommond.TypeFriendTalkInput, jsonData)
	if err != nil {
		fmt.Println("client write err: ", err)
		return err
	}

	return nil
}
