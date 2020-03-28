package rpc

import (
	context "context"
	"myserver/internal/core/netutil/breaker"
	"sync"
	"time"

	grpc "google.golang.org/grpc"
)

var (
	_defaultCliConf = &ClientConfig{
		Dial:              time.Duration(10 * time.Second),
		Timeout:           time.Duration(time.Millisecond * 250),
		Subset:            50,
		KeepAliveInterval: time.Duration(60 * time.Second),
		KeepAliveTimeout:  time.Duration(time.Second * 20),
	}
)

type ClientConfig struct {
	Dial                   time.Duration
	Timeout                time.Duration
	Breaker                *breaker.Config
	Method                 map[string]*ClientConfig
	Clusters               []string
	Zone                   string
	Subset                 int
	NonBlock               bool // 表示在Dial 的时候 是否要用 WithBlock
	KeepAliveInterval      time.Duration
	KeepAliveTimeout       time.Duration
	KeepAliveWithoutStream bool
}

type Client struct {
	conf    *ClientConfig
	breaker *breaker.Group
	mutex   sync.RWMutex

	opts    []grpc.DialOption
	handler []grpc.UnaryClientInterceptor
}

func (c *Client) SetConfig(conf *ClientConfig) error {

	if conf == nil {
		conf = _defaultCliConf
	}

	if conf.Dial == 0 {
		conf.Dial = time.Duration(10 * time.Second)
	}

	if conf.Timeout <= 0 {
		conf.Timeout = time.Duration(time.Millisecond * 20)
	}

	if conf.Subset <= 0 {
		conf.Subset = 50
	}

	if conf.KeepAliveInterval <= 0 {
		conf.KeepAliveInterval = time.Duration(time.Second * 60)
	}

	if conf.KeepAliveTimeout <= 0 {
		conf.KeepAliveTimeout = time.Duration(time.Second * 20)
	}

	c.mutex.Lock()
	c.conf = conf

	if c.breaker == nil {
		c.breaker = breaker.NewGroup(conf.Breaker)
	} else {
		c.breaker.Reload(conf.Breaker)
	}

	c.mutex.Unlock()
	return nil
}

func NewClient(conf *ClientConfig, opt ...grpc.DialOption) *Client {
	c := new(Client)

	if err := c.SetConfig(conf); err != nil {
		panic(err)
	}

	// c.UseOpt(grpc.WithBalancerName("p2c"))
	c.UseOpt(opt...)
	return c
}

func (c *Client) cloneOpts() []grpc.DialOption {
	dialOptions := make([]grpc.DialOption, len(c.opts))
	copy(dialOptions, c.opts)
	return dialOptions
}

func (c *Client) UseOpt(opt ...grpc.DialOption) *Client {
	c.opts = append(c.opts, opt...)
	return c
}

func (c *Client) Use(handler ...grpc.UnaryClientInterceptor) *Client {

	finalSize := len(c.handler) + len(handler)
	mergedHandler := make([]grpc.UnaryClientInterceptor, finalSize)
	copy(mergedHandler, c.handler)
	copy(mergedHandler[:len(c.handler)], handler)
	c.handler = mergedHandler
	return c
}

func (c *Client) Dial(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	opts = append(opts, grpc.WithInsecure())
	return c.dial(ctx, target, opts...)
}

func (c *Client) dial(ctx context.Context, target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	// dialOptions := c.cloneOpts()

	// if !c.conf.NonBlock {
	// 	dialOptions = append(dialOptions, grpc.WithBlock())
	// }

	// dialOptions = append(dialOptions, grpc.WithKeepaliveParams(keepalive.ClientParameters{
	// 	Time:                time.Duration(c.conf.KeepAliveInterval),
	// 	Timeout:             time.Duration(c.conf.KeepAliveTimeout),
	// 	PermitWithoutStream: !c.conf.KeepAliveWithoutStream,
	// }))
	return nil, nil
}
