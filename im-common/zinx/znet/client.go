package znet

import (
	"fmt"
	"go-im/im-common/zinx/iface"
	"net"
	"os"
)

//iServer 接口实现，定义一个Server服务类
type Client struct {
	Name       string								//客户端名称
	LinkAddr   string
	msgHandler iface.IMsgHandle  //消息处理站
	Conn       iface.IConnection //连接
}

/*
  创建一个服务器句柄
*/
func NewClient() iface.IClient {
	c := &Client{
		Name:       "ming's client",
		LinkAddr:   "localhost:8999",
		msgHandler: NewMsgHandle(),
	}
	return c
}

//客户端启动
func (c *Client) Start() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", c.LinkAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println("tcpAddr :")
	conn, err1 := net.DialTCP("tcp", nil, tcpAddr)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err1.Error())
		os.Exit(1)
	}
	c.msgHandler.StartWorkerPool()
	c.Conn = NewClientConnection(conn, c.msgHandler)
	c.Conn.Start()
}



//停止服务
func (c *Client) Stop() {
	fmt.Println("[STOP] Zinx Client , name ", c.Name)
	c.Conn.Stop()
}


//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (c *Client) AddRouter(msgId uint32, router iface.IRouter) {
	c.msgHandler.AddRouter(msgId, router)
}

//得到链接管理
func (c *Client) GetConn() iface.IConnection {
	return c.Conn
}


