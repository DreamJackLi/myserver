package mypaladin

import "context"

type EventType int

const (
	EventAdd EventType = iota
	EventUpdate
	EventRemove
)

type Event struct {
	Event EventType
	Key   string
	Value string
}

type Watcher interface {
	WatchEvent(context.Context, ...string) <-chan Event
	Close() error
}

type Setter interface {
	Set(string) error
}

type Getter interface {
	Get(string) *Value
	GetAll() *Map
}

type Client interface {
	Getter
	Watcher
}
