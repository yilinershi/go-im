package api

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/netCommond"
	"go-im/im-common/proto"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"go-im/im-server/link"
)

type TalkRouter struct {
	znet.BaseRouter
}

func (*TalkRouter) Handle(request iface.IRequest) {
	msg := request.GetData()
	msgId := request.GetMsgID()
	fmt.Printf("get msg from client, msgId=%d,msg=%s\n", msgId, string(msg))

	input := new(proto.TalkMessageInput)
	err := json.Unmarshal(msg, input)
	if err != nil {
		fmt.Println("json unmarshal talk err, err=", err)
		return
	}
	client := link.Clients.Get(input.To)
	var resp = &proto.TalkMessageSync{
		From:    input.From,
		Content: input.Content,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json marshal TalkRouter err,err=", err)
		return
	}
	client.SendMsg(netCommond.TypeFriendTalkSync, jsonData)
	fmt.Printf("聊天sync, msgId=%d,msg=%s\n", netCommond.TypeFriendTalkSync, string(jsonData))
}
