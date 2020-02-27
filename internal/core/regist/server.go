package regist

import (
	"fmt"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type InitFunc func(*grpc.Server)

var registerMap map[string]*grpc.Server
var registInit map[string]InitFunc
var lock *sync.Mutex

func init() {
	registerMap = make(map[string]*grpc.Server)
	registInit = make(map[string]InitFunc)
	lock = &sync.Mutex{}
}

func InitServer(name string, r InitFunc) {
	fmt.Println("InitServer", name)
	registInit[name] = r
}

func GetInitServer(name string) InitFunc {
	f, ok := registInit[name]

	if ok {
		return f
	} else {
		return nil
	}

}

func RegistServer(name string, endPort string, f InitFunc, opt ...grpc.ServerOption) error {
	var s *grpc.Server
	if opt != nil {
		s = grpc.NewServer(opt...)
	} else {
		s = grpc.NewServer()
	}
	reflection.Register(s)
	lock.Lock()
	registerMap[name] = s
	lock.Unlock()
	f(s)
	lis, err := net.Listen("tcp", endPort)

	if err != nil {
		return err
	}

	if err = s.Serve(lis); err != nil {
		fmt.Println("RegistServer error is ", err)
		return err
	}
	fmt.Println("RpcServer Exit Name is ", name)
	return nil

}

func CloseServer() {
	lock.Lock()
	for k, v := range registerMap {
		temp := v
		temp.Stop()
		fmt.Println("close grpc.Server name is ", k)
	}
	lock.Unlock()
}
