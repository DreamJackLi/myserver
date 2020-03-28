package conn

import (
	"fmt"
	"myserver/api/base"
	"myserver/config"
	"myserver/internal/core/message"
	"myserver/internal/gateway/server/router"
	"myserver/tool/uid"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

var c *Conn

// 一个用户可能有多个链接 这时候区分就用GroupID
type Conn struct {
	ConnID     string
	GroupID    string
	ServerID   string
	ForwordFor string
	isStop     chan struct{}
	wg         *sync.WaitGroup
	wbConn     *websocket.Conn

	readChan       chan []byte
	writeBuffer    chan *message.Message
	writeBuffer_v2 chan *base.Msg

	// Heart
	CheckTime time.Duration
	// 0 表示未开启 1表示开启 2表示关闭
	statues int32

	HeartTime int64
}

func NewConn(wg *sync.WaitGroup, wbConn *websocket.Conn, forwardFor string) *Conn {
	connID := uid.GetStringUUID("")
	groupID := uid.GetStringUUID("")
	return &Conn{
		wg:         wg,
		ConnID:     connID,
		GroupID:    groupID,
		wbConn:     wbConn,
		ForwordFor: forwardFor,
		// 这里后面可以把wbConn中读到的message放进来 增加吞吐
		readChan:       make(chan []byte, 1024),
		writeBuffer:    make(chan *message.Message, 1024),
		writeBuffer_v2: make(chan *base.Msg, 1024),
		isStop:         make(chan struct{}),
		// 这里后面加心跳
		CheckTime: 5,
		statues:   0,
		HeartTime: 0,
	}
}

func (c *Conn) Start() {
	go c.workProcc()
	go c.handMsg()
	c.wg.Add(2)
	atomic.AddInt32(&c.statues, 1)
	fmt.Println("Conn is Start ConnID is ", c.ConnID)
}

func (c *Conn) Close() {
	if atomic.LoadInt32(&c.statues) == 2 {
		return
	}
	fmt.Println("Close ConnID is ", c.ConnID)
	atomic.AddInt32(&c.statues, 1)
	c.wbConn.Close()
	close(c.writeBuffer)
	close(c.writeBuffer_v2)
	close(c.readChan)
	close(c.isStop)
}

func (c *Conn) handMsg() {
	for {
		select {
		case <-c.isStop:
			c.wg.Done()
			return
		case m, ok := <-c.readChan:
			if ok {
				c.readMsg_v2(m)
			}

		}
	}
}

func (c *Conn) workProcc() {
	heart := time.NewTicker(time.Duration(config.Get().ConnCfg.CheckTime) * time.Second)
	for {
		select {
		case <-c.isStop:
			heart.Stop()
			c.wg.Done()
			return
		case <-heart.C:

		case m, ok := <-c.writeBuffer_v2:
			if ok {
				c.writeMsg_v2(m)
			} else {
				c.Close()
				DeleteConnByID(c)
				c.wg.Done()
				return
			}
		default:
			_, b, err := c.wbConn.ReadMessage()
			if err != nil {
				fmt.Println("ReadMessage Error ", err)
				c.Close()
				DeleteConnByID(c)
				c.wg.Done()
				return
			}
			c.readChan <- b
			// c.readMsg_v2(b)
		}

	}
}

func (c *Conn) readMsg_v2(b []byte) {
	m, err := message.UnPacker_v2(b)
	if err != nil {
		fmt.Println("Read Message Fail err ", err)
		return
	}
	fmt.Println("Read Message ApiID is ", m.ApiId)

	f := router.GetHandler_v2(int32(m.ApiId))
	if f == nil {
		fmt.Println("No Regist Handler ApiID is ", m.ApiId)
		return
	} else {
		res := f(m)
		if res != nil {
			c.WriteMessage_v2(res)
		}
	}

}

// func (c *Conn) readMsg(b []byte) {
//   m := message.UnPacker(b)
//   if m == nil {
//     fmt.Println("Read Message Fail ")
//     return
//   }
//   fmt.Println("Read Message ApiID is ", m.Head.ApiID)

//   f := router.GetHandler(m.Head.ApiID)
//   if f == nil {
//     fmt.Println("No Regist Handler ApiID is ", m.Head.ApiID)
//     return
//   } else {

//     res := f(m)
//     if res != nil {
//       c.WriteMessage(res)
//     }
//   }

// }

// func (c *Conn) writeMsg(m *message.Message) {
//   b := message.Pack(m)
//   if b == nil {
//     fmt.Println("Write Message Error ApiID is ", m.Head.ApiID)
//     return
//   }
//   err := c.wbConn.WriteMessage(websocket.BinaryMessage, b)

//   if err != nil {
//     fmt.Println("213123")
//   }

// }

func (c *Conn) writeMsg_v2(m *base.Msg) {
	fmt.Println("Write Msg apiID is ", m.ApiId)
	b, err := message.Pack_v2(m)
	if err != nil {
		fmt.Printf("Write Message Error is %s , ApiID is %d\n", err.Error(), m.ApiId)
		return
	}
	err = c.wbConn.WriteMessage(websocket.BinaryMessage, b)

	if err != nil {
		fmt.Println("wbConn WriteMessage")
	}

}

// func (c *Conn) write() {
//   for {
//     select {
//     case <-c.isStop:
//       c.wg.Done()
//       return
//     case m, ok := <-c.writeBuffer:
//       if ok {
//         fmt.Println("Write Message ApiID is ", m.Head.ApiID)
//         c.writeMsg(m)
//       }
//     }
//   }
// }

func (c *Conn) WriteMessage(m *message.Message) {
	select {
	case <-c.isStop:
		return
	case c.writeBuffer <- m:
	}
}

func (c *Conn) WriteMessage_v2(m *base.Msg) {
	select {
	case <-c.isStop:
		return
	case c.writeBuffer_v2 <- m:
	}
}
