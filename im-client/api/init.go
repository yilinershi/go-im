package api

import (
	"fmt"
	"go-im/im-common/zinx/iface"
)

var conn iface.IConnection

func Init(clientConn iface.IConnection)  {
	conn =clientConn
	fmt.Println("client init conn success!")
}