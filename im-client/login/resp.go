package login

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/eventCenter"
	"go-im/im-common/proto"
	"go-im/im-common/zinx/iface"
	"go-im/im-common/zinx/znet"
	"reflect"
)

//持有底层router，方便client注册
type Router struct {
	znet.BaseRouter
}

//服务器响应的回调
func (*Router) Handle(request iface.IRequest) {
	msg := request.GetData()
	fmt.Println("get msg from server, msg=", string(msg))
	var resp = new(proto.LoginResp)
	err := json.Unmarshal(msg, resp)
	if err != nil {
		fmt.Println("json Unmarshal resp err, err=", err)
		return
	}

	if reflect.DeepEqual(resp.Error, errorCode.LoginPasswordError) {
		fmt.Println("登录失败 >> err=", resp.Error)
		TipLogin(nil)
		return
	}

	if reflect.DeepEqual(resp.Error, errorCode.LoginAccountNotExist) {
		fmt.Println("登录失败 >> err=", resp.Error)
		eventCenter.Emit(eventCenter.OnTipRegister, nil)
		return
	}

	eventCenter.Emit(eventCenter.OnLoginSuccess, nil)

}
