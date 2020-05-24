package proto

import "go-im/im-common/errorCode"

type RegisterReq struct {
	Account   string
	Password1 string
	Password2 string
	Age     int
}

type RegisterResp struct {
	Error	*errorCode.ErrorBase
}



