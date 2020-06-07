package link

import "go-im/im-common/zinx/iface"

type Client struct {
	Account string
	Conn    iface.IConnection
}

//客户端下线
func (this *Client) LostConnection() {

}

//未来需要将这里由protobuf转为[]byte发送
func (this *Client) SendMsg(msgId uint32, data []byte) {
	this.Conn.SendMsg(msgId, data)
}
