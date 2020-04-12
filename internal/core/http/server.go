package http

import (
	"fmt"
	"myserver/internal/core/ip"
	"regexp"
	"sync"
	"sync/atomic"
	"time"
)

type Handler interface {
	ServerHttp(c *Context)
}

type HandlerFunc func(*Context)

// ServerConfig ServerConfig
type ServerConfig struct {
	Network      string        `dns:"network"`
	Addr         string        `dns:"address"`
	Timeout      time.Duration `dns:"query.timeout"`
	ReadTimeout  time.Duration `dns:"query.readTimeout"`
	WriteTimeout time.Duration `dns:"query.writeTimeout"`
}

type MethodConfig struct {
	Timeout time.Duration
}

// Engine Engine
type Engine struct {
	RouterGroup

	lock sync.RWMutex
	conf *ServerConfig

	address string

	trees     methodTrees
	server    atomic.Value
	metastore map[string]map[string]interface{}

	pcLock        sync.RWMutex
	methodConfigs map[string]*MethodConfig

	injections []injection

	UserRawPath bool

	UnescapePathValues bool

	HandlerMethodNotAllowed bool

	allNoRoute  []HandlerFunc
	allNoMethod []HandlerFunc
	noRouter    []HandlerFunc
	noMethod    []HandlerFunc
}

type injection struct {
	pattern  *regexp.Regexp
	handlers []HandlerFunc
}

// NewServer NewServer
func NewServer(conf *ServerConfig) *Engine {

	if conf == nil {

	}

	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
			root:     true,
		},
		address:                 ip.InternalIp(),
		trees:                   make(methodTrees, 0, 9),
		metastore:               make(map[string]map[string]interface{}),
		methodConfigs:           make(map[string]*MethodConfig),
		HandlerMethodNotAllowed: true,
		injections:              make([]injection, 0),
	}

	if err := engine.SetConfig(conf); err != nil {
		panic(err)
	}

	engine.RouterGroup.engine = engine

	return engine
}

func (engine *Engine) SetConfig(conf *ServerConfig) (err error) {
	if conf.Timeout <= 0 {
		return fmt.Errorf("config timeout must greater than 0")
	}

	if conf.Network == "" {
		conf.Network = "tcp"
	}

	engine.lock.Lock()
	engine.conf = conf
	engine.lock.Unlock()
	return nil
}

func (engine *Engine) addRoute(method, path string, handlers ...HandlerFunc) {

	if path[0] != '/' {
		panic("path must begin with '/' ")
	}

	if method == "" {
		panic("http mathod can not be empty")
	}

	if len(handlers) == 0 {
		panic("there must be at least on handler")
	}

	if _, ok := engine.metastore[path]; !ok {

	}
}
