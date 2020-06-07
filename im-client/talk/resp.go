package talk

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/proto"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
)

//持有底层router，方便client注册
type Router struct {
	znet.BaseRouter
}

//服务器响应的回调
func (*Router) Handle(request iface.IRequest) {
	msg := request.GetData()
	fmt.Println("friend resp, get msg from server, msg=", string(msg))
	var resp = new(proto.TalkMessageSync)
	err := json.Unmarshal(msg, resp)
	if err != nil {
		fmt.Println("json Unmarshal resp err, err=", err)
		return
	}

	//获取好友列表后，显示选择好友并对话的提示框
	onSyncFriendTalk(resp.From, resp.Content)
}
