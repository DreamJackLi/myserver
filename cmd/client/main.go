package main

import (
	"fmt"
	"myserver/api/apiemployee"
	"myserver/api/base"
	"myserver/internal/core/message"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var servers []string = []string{"ws://errand.kuaixg.cn/ws"}

//var servers []string = []string{"ws://47.106.222.37:8082/ws"}
var total int32

func main() {
	u := url.URL{Scheme: "ws", Host: "192.168.3.38:8080", Path: "/begin"}
	// u := url.URL{Scheme: "ws", Host: "errand.kuaixg.cn", Path: "/begin"}
	fmt.Println(u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	wg := sync.WaitGroup{}
	if err != nil {
		fmt.Println("Dial Error ", err)
		return
	}

	msg := message.NewMessage_v2(base.Api_BuildConn, nil)

	b, _ := msg.Marshal()

	c.WriteMessage(websocket.BinaryMessage, b)

	go func() {

		for {
			_, b, err := c.ReadMessage()

			if err != nil {
				fmt.Println("ReadMessage Error ", err)
				return
			}

			m, err := message.UnPacker_v2(b)
			if err != nil {
				fmt.Println("UnPacker Err ", err)
				continue
			}
			fmt.Println(m.ApiId)
			if m.ApiId == base.Api_BuildConn {
				out := &base.Head{}

				err = out.Unmarshal(m.Data)
				fmt.Println("Get ApiId is ConnID is ", m.ApiId, out.ConnID)

				if err != nil {
					fmt.Println("Unmarshal 1111 Error ", err)
					continue
				}

			} else if m.ApiId == base.Api_WebLogin {

				out := &apiemployee.WebLoginRes{}

				err = out.Unmarshal(m.Data)

				fmt.Println(out.Info)

				if err != nil {
					fmt.Println("Unmarshal 2222 Error ", err)
					continue
				}

			}

		}
	}()
	wg.Add(1)

	go func() {

		tick := time.NewTicker(2 * time.Second)

		for {
			select {
			case <-tick.C:

				in := &apiemployee.WebLoginReq{
					Head: &base.Head{
						ConnID: "1212131",
					},
					UserName: "xiaoka2020",
					Pwd:      "xiaoka",
				}
				inB, _ := in.Marshal()
				m := message.NewMessage_v2(base.Api_WebLogin, inB)
				res, err := message.Pack_v2(m)
				if err != nil {
					fmt.Println("Msg Pack Err ", err)
					continue
				}
				c.WriteMessage(websocket.BinaryMessage, res)

			}
		}

	}()
	wg.Add(1)
	wg.Wait()
}
