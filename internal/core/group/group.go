package group

import (
	"robot-server/internal/core/conn"
	"sync"
)

var g *Group

type Group struct {
	locker   *sync.Mutex
	groupMap map[string]*conn.Conn
}

func newGroup() *Group {
	return &Group{}
}

func init() {
	g = newGroup()
}

func GetConnByGroup(groupID string) *conn.Conn {
	return nil
}
