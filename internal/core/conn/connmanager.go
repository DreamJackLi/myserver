package conn

import (
	"fmt"
	"sync"
)

var cm *ConnManager

func init() {
	cm = newConnManager()
}

type ConnManager struct {
	lock    *sync.Mutex
	connMap map[string]*Conn
	num     int32
}

func newConnManager() *ConnManager {
	return &ConnManager{
		lock:    &sync.Mutex{},
		connMap: make(map[string]*Conn),
		num:     0,
	}
}

func GetConnByConn(connID string) *Conn {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	c, ok := cm.connMap[connID]

	if ok {
		return c
	} else {
		return nil
	}
}

func SetConnByID(c *Conn) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.connMap[c.ConnID] = c
	cm.num++
	fmt.Println("SetConn Cur Conn Num is ", cm.num)
}

func DeleteConnByID(c *Conn) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	_, ok := cm.connMap[c.ConnID]
	if ok {
		delete(cm.connMap, c.ConnID)
	}
	cm.num--
	fmt.Println("Delete Conn Cur Conn Num is ", cm.num)
}

func CloseConnManager() {
	cm.lock.Lock()
	for _, v := range cm.connMap {
		v.Close()
		delete(cm.connMap, v.ConnID)
	}
	cm.lock.Unlock()
}
