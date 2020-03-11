package rpc

import (
	context "context"
	fmt "fmt"
	"testing"
	"time"

	grpc "google.golang.org/grpc"
)

const (
	_testAddr = "127.0.0.1:9090"
)

type HelloS struct {
}

func (h *HelloS) HelloWrold(ctx context.Context, req *HelloWroldReq) (*HelloWroldRes, error) {
	fmt.Println("HelloWrold 11111")
	panic("qweqwe")
	// return &HelloWroldRes{}, nil
}

func runServer() {
	s := NewServer(&ServerCfg{Addr: _testAddr, TimeOut: time.Duration(time.Second)})
	RegisterHelloServer(s.Server(), &HelloS{})
	s.Use()
	if _, err := s.Start(); err != nil {
		fmt.Println(err)
	}
}

func runClient() {
	c, err := grpc.Dial(_testAddr, grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("RunClient Err%v\n : ", err))
	}

	hC := NewHelloClient(c)

	res, err := hC.HelloWrold(context.Background(), &HelloWroldReq{Name: 12})

	fmt.Println(res)
}

func TestRpc(t *testing.T) {
	runServer()

	time.Sleep(3 * time.Second)
	runClient()

	select {}
}
