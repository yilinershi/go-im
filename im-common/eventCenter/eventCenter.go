package eventCenter

var eventMap = make(map[string][]func(interface{}))

func On(name string, f func(interface{})) {
	list := eventMap[name]
	list = append(list, f)
	eventMap[name] = list
}

func Off(name string) {
	eventMap[name] = nil
}

func Emit(name string, pram interface{}) {
	list := eventMap[name]
	for _, f := range list {
		f(pram)
	}
}
