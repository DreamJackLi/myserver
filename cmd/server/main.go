package main

import (
	"fmt"
	"myserver/config"
	"os"

	"myserver/internal/core/regist"
	"myserver/internal/core/sig"
	"myserver/internal/gateway/server"
	_ "myserver/internal/micro/employ"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	registServer(wg)
	// err := coredb.Init()
	// if err != nil {
	//   fmt.Println("cmd main coredb.Init Error ", err)
	//   return
	// }
	wg.Add(1)
	go sig.Signal(func(os.Signal) {
		fmt.Println("Close signal Come ")
		server.Ser.Close()
		regist.CloseClientConn()
		regist.CloseServer()
		wg.Done()
	})

	wg.Add(1)
	server.Start(wg)

	// test close server
	// wg.Add(1)
	// go func() {
	//   t := time.NewTimer(5 * time.Second)

	//   select {
	//   case <-t.C:

	//     fmt.Println("timer is Up ")
	//     server.Ser.Close()
	//     regist.CloseClientConn()
	//     regist.CloseServer()
	//     wg.Done()
	//   }
	// }()
	wg.Wait()

}

func registServer(wg *sync.WaitGroup) {
	for _, v := range regist.ServerName {
		f := regist.GetInitServer(v)

		if f != nil {
			wg.Add(1)
			go func(name string, fun regist.InitFunc) {
				config.LoadService(name)
				addr := fmt.Sprintf(":%d", config.Get().ServerCfg[name].Port)

				fmt.Printf("RegistServer Name is%s , addr is %s \n ", name, addr)
				regist.RegistServer(name, addr, fun)
				wg.Done()
			}(v, f)
		}
	}

}
