package rpc

import (
	"context"
	"fmt"
	"math"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var (
	_defaultSerConf = &ServerCfg{
		Network:           "tcp",
		Addr:              "0.0.0.0:9000",
		TimeOut:           time.Duration(time.Second),
		IdleTimeOut:       time.Duration(180 * time.Second),
		MaxLifeTime:       time.Duration(2 * time.Hour),
		ForecdCloseWait:   time.Duration(20 * time.Second),
		KeepAliveInterval: time.Duration(60 * time.Second),
		KeepAliveTimeOut:  time.Duration(20 * time.Second),
	}
	_abortIndex int8 = math.MaxInt8 / 2
)

type ServerCfg struct {
	Network           string
	Addr              string
	TimeOut           time.Duration
	IdleTimeOut       time.Duration
	MaxLifeTime       time.Duration
	ForecdCloseWait   time.Duration
	KeepAliveInterval time.Duration
	KeepAliveTimeOut  time.Duration
	LogFlag           int8
}

type Server struct {
	conf  *ServerCfg
	mutex sync.RWMutex

	server   *grpc.Server
	handlers []grpc.UnaryServerInterceptor
}

func (s *Server) SetConfig(conf *ServerCfg) (err error) {
	if conf == nil {
		conf = _defaultSerConf
	}

	if conf.TimeOut <= 0 {
		conf.TimeOut = time.Duration(time.Second)
	}

	if conf.IdleTimeOut <= 0 {
		conf.TimeOut = time.Duration(60 * time.Second)
	}

	if conf.MaxLifeTime <= 0 {
		conf.TimeOut = time.Duration(2 * time.Hour)
	}

	if conf.ForecdCloseWait <= 0 {
		conf.TimeOut = time.Duration(20 * time.Second)
	}

	if conf.KeepAliveInterval <= 0 {
		conf.TimeOut = time.Duration(60 * time.Second)
	}

	if conf.KeepAliveTimeOut <= 0 {
		conf.TimeOut = time.Duration(20 * time.Second)
	}

	if conf.Addr == "" {
		conf.Addr = "0.0.0.0:9000"
	}

	if conf.Network == "" {
		conf.Network = "tcp"
	}

	s.mutex.Lock()
	s.conf = conf
	s.mutex.Unlock()

	return nil
}

func (s *Server) interceptor(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	var (
		i     int
		chain grpc.UnaryHandler
	)

	n := len(s.handlers)

	if n == 0 {
		handler(ctx, req)
	}

	chain = func(ic context.Context, ir interface{}) (interface{}, error) {
		fmt.Println("chain ", i)
		if i == n-1 {
			return handler(ic, ir)
		}
		i++
		return s.handlers[i](ic, ir, args, chain)
	}

	return s.handlers[0](ctx, req, args, chain)

}

func (s *Server) Use(handler ...grpc.UnaryServerInterceptor) *Server {
	finalSize := len(s.handlers) + len(handler)
	if finalSize >= int(_abortIndex) {
		panic("warden : server use too many handlers")
	}
	mergeHandlers := make([]grpc.UnaryServerInterceptor, finalSize)
	copy(mergeHandlers, s.handlers)
	copy(mergeHandlers[len(s.handlers):], handler)

	s.handlers = mergeHandlers
	return s
}

func (s *Server) recovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if rerr := recover(); rerr != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				rs := runtime.Stack(buf, false)

				if rs > size {
					rs = size
				}

				buf = buf[:rs]
				pl := fmt.Sprintf("grpc server panic: %v\n%v\n%s\n", req, rerr, buf)

				fmt.Fprintf(os.Stderr, pl)
				err = status.Errorf(codes.Unknown, "-504")
			}
		}()
		fmt.Println("recovery ", args.FullMethod, args.Server)
		resp, err = handler(ctx, req)
		return
	}
}

func (s *Server) handle() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("handle ", args.FullMethod, args.Server)
		resp, err = handler(ctx, req)
		return
	}
}

func (s *Server) Start() (*Server, error) {
	_, err := s.startWithAddr()
	if err != nil {
		return nil, err
	}
	fmt.Println("Server Start")
	return s, nil
}

func (s *Server) Server() *grpc.Server {
	return s.server
}

func (s *Server) startWithAddr() (net.Addr, error) {
	lis, err := net.Listen(s.conf.Network, s.conf.Addr)

	if err != nil {
		return nil, err
	}

	reflection.Register(s.server)

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return lis.Addr(), nil
}

func (s *Server) Serve(lis net.Listener) error {
	return s.server.Serve(lis)
}

func NewServer(conf *ServerCfg, opt ...grpc.ServerOption) *Server {
	s := &Server{}
	if err := s.SetConfig(conf); err != nil {
		panic(errors.Errorf("warden: set config failed! err: %s", err.Error()))
	}

	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(s.conf.IdleTimeOut),
		MaxConnectionAgeGrace: time.Duration(s.conf.ForecdCloseWait),
		Time:                  time.Duration(s.conf.KeepAliveInterval),
		Timeout:               time.Duration(s.conf.KeepAliveTimeOut),
		// MaxConnectionAgeGrace: time.Duration(s.conf.MaxLifeTime),
	})
	// 这里注册interceptor
	// 由这里的 interceptor 去调用server 里 handler 中的各种拦截器
	opt = append(opt, keepParam, grpc.UnaryInterceptor(s.interceptor))

	s.server = grpc.NewServer(opt...)
	s.Use(s.recovery(), s.handle())
	return s
}

func (s *Server) ShutDown(ctx context.Context) (err error) {
	ch := make(chan struct{})
	go func() {
		s.server.GracefulStop()
		close(ch)
	}()
	select {
	case <-ctx.Done():
		s.server.Stop()
		err = ctx.Err()
	case <-ch:
	}
	return
}
