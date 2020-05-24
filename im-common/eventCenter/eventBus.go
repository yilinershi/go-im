package eventCenter

import (
	"unsafe"
)

type Dispatcher struct {
	listeners map[string]*EventChain
}

type EventChain struct {
	chs       []chan *Event
	callbacks []*EventCallback
}

func createEventChain() *EventChain {
	return &EventChain{chs: []chan *Event{}, callbacks: []*EventCallback{}}
}

type Event struct {
	eventName string
	Params    map[string]interface{}
}

func CreateEvent(eventName string, params map[string]interface{}) *Event {
	return &Event{eventName: eventName, Params: params}
}

type EventCallback func(*Event)

var _instance *Dispatcher

func SharedDispatcher() *Dispatcher {
	if _instance == nil {
		_instance = &Dispatcher{}
		_instance.Init()
	}

	return _instance
}

func (this *Dispatcher) Init() {
	this.listeners = make(map[string]*EventChain)
}

func (this *Dispatcher) AddEventListener(eventName string, callback *EventCallback) {
	eventChain, ok := this.listeners[eventName]
	if !ok {
		eventChain = createEventChain()
		this.listeners[eventName] = eventChain
	}

	exist := false
	for _, item := range eventChain.callbacks {
		a := *(*int)(unsafe.Pointer(item))
		b := *(*int)(unsafe.Pointer(callback))
		if a == b {
			exist = true
			break
		}
	}

	if exist {
		return
	}

	ch := make(chan *Event)

	eventChain.chs = append(eventChain.chs[:], ch)
	eventChain.callbacks = append(eventChain.callbacks[:], callback)

	go this.handler(eventName, ch, callback)
}

func (this *Dispatcher) handler(eventName string, ch chan *Event, callback *EventCallback) {
	for {
		event := <-ch
		if event == nil {
			break
		}
		go (*callback)(event)
	}
}

func (this *Dispatcher) RemoveEventListener(eventName string, callback *EventCallback) {
	eventChain, ok := this.listeners[eventName]
	if !ok {
		return
	}

	var ch chan *Event
	exist := false
	key := 0
	for k, item := range eventChain.callbacks {
		a := *(*int)(unsafe.Pointer(item))
		b := *(*int)(unsafe.Pointer(callback))
		if a == b {
			exist = true
			ch = eventChain.chs[k]
			key = k
			break
		}
	}

	if exist {
		ch <- nil
		eventChain.chs = append(eventChain.chs[:key], eventChain.chs[key+1:]...)
		eventChain.callbacks = append(eventChain.callbacks[:key], eventChain.callbacks[key+1:]...)
	}
}

func (this *Dispatcher) DispatchEvent(event *Event) {
	eventChain, ok := this.listeners[event.eventName]
	if ok {
		for _, chEvent := range eventChain.chs {
			chEvent <- event
		}
	}
}
