package register

import (
	"encoding/json"
	"fmt"
	"go-im/im-common/errorCode"
	"go-im/im-common/eventCenter"
	"go-im/im-common/model"
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
	var resp = new(model.RegisterResp)
	err := json.Unmarshal(msg, resp)
	if err != nil {
		fmt.Println("json Unmarshal resp err, err=", err)
		return
	}

	if !reflect.DeepEqual(resp.Error, errorCode.OK) {
		fmt.Println(resp.Error.Desc)
		TipRegister(nil)
		return
	}

	fmt.Println("注册成功,请进行登录")
	eventCenter.Emit(eventCenter.OnTipLogin, nil)
}
