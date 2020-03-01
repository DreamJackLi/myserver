package server

import (
	"fmt"
	"myserver/api/base"
	"myserver/internal/core/conn"
	"myserver/internal/core/message"
	"net/http"
)

func buildConn(w http.ResponseWriter, r *http.Request) {

	forward := r.Header.Get("X-FORWARDED-FOR")
	if forward == "" {
		forward = r.RemoteAddr
	}

	fmt.Println("build Conn come client ip is ", forward)
	wbConn, err := Ser.upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Upgrade Error ", err)
		return
	}

	_, b, err := wbConn.ReadMessage()

	if err != nil {
		fmt.Println("Bad WebSocket")
		return
	}

	m, err := message.UnPacker_v2(b)

	if err != nil {
		fmt.Println("BuildConn Message illegal")
		return
	}

	if m.ApiId != base.Api_BuildConn {
		fmt.Println("BuildConn ApiID illegal")
		return
	}

	c := conn.NewConn(&Ser.wg, wbConn, forward)

	conn.SetConnByID(c)

	c.Start()

	head := &base.Head{
		ConnID:   c.ConnID,
		GroupID:  c.GroupID,
		AppKey:   "qweqwe",
		ServerID: Ser.ServerID,
	}

	// m := &message.Message{
	//   Head: head,
	//   Body: []byte("asdasd"),
	// }
	hb, _ := head.Marshal()
	bm := &base.Msg{
		ApiId: base.Api_BuildConn,
		Data:  hb,
	}

	// c.WriteMessage(m)
	c.WriteMessage_v2(bm)
}
