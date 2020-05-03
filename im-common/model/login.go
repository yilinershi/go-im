
//因为数据足够小，model层供client与server共用，这里放到common里，方便两端使用
package model

import "go-im/im-common/errorCode"

type LoginReq struct {
	Account  string
	Password string
}

type LoginResp struct {
	Error	*errorCode.ErrorBase
}
