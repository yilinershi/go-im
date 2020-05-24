package eventCenter

//事件中心，该事件心中只支持一事件一响应
var eventMap = make(map[string][]func(interface{}))

//监听事件
func On(name string, f func(interface{})) {
	list := eventMap[name]
	list = append(list, f)
	eventMap[name] = list
}


//注销监听
func Off(name string) {
	eventMap[name] = nil
}

//提交事件
func Emit(name string, pram interface{}) {
	list := eventMap[name]
	for _, f := range list {
		f(pram)
	}
}
