package proto

import (
	"go-im/im-common/errorCode"
	"go-im/im-common/model"
)


//获取好友列表req
type FriendListReq struct {
	Account  string  //当前连接的用户
}

//获取好友列表resp
type FriendListResp struct {
	Error   *errorCode.ErrorBase
	Friends []model.FriendInfo
}
