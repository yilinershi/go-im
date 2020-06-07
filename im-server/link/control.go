package link

import (
	"sync"
)

type linkClientControler struct {
	c     map[string]*Client
	pLock sync.RWMutex //保护Players的互斥读写机制
}

var Clients *linkClientControler

func init() {
	Clients = &linkClientControler{
		c: make(map[string]*Client),
	}
}

/*
* 将客户端的连接加入到所有连接的列表，只有当客户端登录成功后才可以加入
*/
func (this *linkClientControler) Add(u *Client) {
	this.pLock.Lock()
	this.c[u.Account] = u
	this.pLock.Unlock()
}

func (this *linkClientControler) Remove(account string) {
	this.pLock.Lock()
	delete(this.c, account)
	this.pLock.Unlock()
}

func (this *linkClientControler) Get(account string) *Client {
	this.pLock.RLock()
	defer this.pLock.RUnlock()
	return this.c[account]
}

func (this *linkClientControler) GetAll(account string) []*Client {
	this.pLock.RLock()
	defer this.pLock.RUnlock()
	target := make([]*Client, 0)
	for _, v := range this.c {
		target = append(target, v)
	}
	return target
}
