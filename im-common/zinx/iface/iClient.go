package iface

//定义服务器接口
type IClient interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router IRouter)
	//得到链接管理
	GetConn() IConnection
}
