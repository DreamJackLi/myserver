package mypaladin

import (
	"strings"
	"sync/atomic"
)

type Map struct {
	values atomic.Value
}

func KeyNamed(key string) string {
	return strings.ToLower(key)
}

// 这里是浅拷贝
func (m *Map) Store(values map[string]*Value) {
	dst := make(map[string]*Value, len(values))

	for k, v := range values {
		dst[KeyNamed(k)] = v
	}
	m.values.Store(dst)
}

// 这里是浅拷贝
func (m *Map) Load() map[string]*Value {
	src := m.values.Load().(map[string]*Value)
	dst := make(map[string]*Value, len(src))

	for k, v := range src {
		dst[k] = v
	}

	return dst
}

func (m *Map) Exist(key string) bool {
	_, ok := m.Load()[KeyNamed(key)]
	return ok
}

func (m *Map) Get(key string) *Value {
	v, ok := m.Load()[KeyNamed(key)]
	if ok {
		return v
	}
	return &Value{}
}

func (m *Map) Keys() []string {
	values := m.Load()
	keys := make([]string, 0, len(values))

	for key := range values {
		keys = append(keys, key)
	}

	return keys
}
