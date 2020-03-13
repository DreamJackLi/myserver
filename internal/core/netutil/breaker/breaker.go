package breaker

import (
	"sync"
	"time"
)

var (
	_mu   sync.RWMutex
	_conf = &Config{
		Window:  time.Duration(3 * time.Second),
		Bucket:  10,
		Request: 100,

		// Percentage of failures must be lower than 33.33%
		K: 1.5,
	}
)

// 熔断器的配置
type Config struct {
	SwitchOff bool // 是否打开熔断器 默认关闭

	// 熔断系数
	K float64

	Window  time.Duration
	Bucket  int
	Request int64
}

func (conf *Config) fix() {
	if conf.K == 0 {
		conf.K = 1.5
	}

	if conf.Request == 0 {
		conf.Request = 100
	}

	if conf.Bucket == 0 {
		conf.Bucket = 10
	}

	if conf.Window == 0 {
		conf.Window = time.Duration(3 * time.Second)
	}

}

type Breaker interface {
	Allow() error
	MarkSuccess()
	MarkFailed()
}

type Group struct {
	mu   sync.RWMutex
	brks map[string]Breaker
	conf *Config
}

func NewGroup(conf *Config) *Group {
	if conf == nil {
		_mu.Lock()
		conf = _conf
		_mu.Unlock()
	} else {
		conf.fix()
	}

	return &Group{
		conf: conf,
		brks: make(map[string]Breaker),
	}
}

func (g *Group) Reload(conf *Config) {
	if conf == nil {
		return
	}

	conf.fix()

	g.mu.Unlock()
	g.conf = conf
	g.brks = make(map[string]Breaker, len(g.brks))
	g.mu.Unlock()

}
