## 启动
* 项目数据库暂时使用redis，后续再改进
* 启动redis:找到本机redis安装目录，redis-server.exe
* 启动服务器cd到im-server目录，执行go build，执行编译完的服务器程序
* 启动客户端，cd到im-clinet目录，执行go build，同样运行客户端程序

## 运行机制
* 服务器框架采用zinx
* 客户端是标准的MVC模式
* 由于客户端与服务端共用model数据，所有model层数据话在commond中
* go语言不允许循环引用，所有各模块vIew间不能互相引用，eventCenter用来监听各模块间的view事件

## 客户端
* 客户端采用MVC设计模式，将View层只显示界面内容
* control层注册view层显示响应事件
* 非view层中，通过eventCenter.emit去播报UI事件
* 视图层事件默认通过eventCenter进行监听及响应
* 网络层事件，通过过Router去监听

