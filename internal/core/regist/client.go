package regist

import (
	"fmt"
	"log"
	"robot-server/config"
	"sync"

	"google.golang.org/grpc"
)

var dis *Discovery

type Discovery struct {
	lock   *sync.Mutex
	serCli map[string]*grpc.ClientConn
}

func newDiscovery() *Discovery {
	return &Discovery{
		lock:   &sync.Mutex{},
		serCli: make(map[string]*grpc.ClientConn),
	}
}

func init() {
	dis = newDiscovery()
}

func registRpcCli(name string) *grpc.ClientConn {
	addr := fmt.Sprintf(":%d", config.Get().ServerCfg[name].Port)
	fmt.Println(addr)
	conn, err := grpc.Dial(addr, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
	}...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil
	}
	return conn
}

func GetCliConn(name string) *grpc.ClientConn {
	dis.lock.Lock()
	defer dis.lock.Unlock()
	c, ok := dis.serCli[name]
	if ok {
		return c
	} else {
		fmt.Println("registRpcCli")
		cli := registRpcCli(name)
		if cli == nil {
			fmt.Println("registRpcCli Fail ServerName is ", name)
		}
		dis.serCli[name] = cli
		return cli
	}
}

func SetCliConn(name string, c *grpc.ClientConn) {
	dis.lock.Lock()
	defer dis.lock.Unlock()
	_, ok := dis.serCli[name]
	if !ok {
		dis.serCli[name] = c
	}
}

func CloseClientConn() {
	dis.lock.Lock()
	for k, v := range dis.serCli {
		temp := v
		temp.Close()
		// v.Close()
		fmt.Println("Close ClientConn name is ", k)
	}
	dis.lock.Unlock()
}
