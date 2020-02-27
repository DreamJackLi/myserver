package server

import (
	"fmt"
	"net/http"
	"robot-server/internal/core/conn"
	"robot-server/internal/gateway/server/router"
	"robot-server/tool/uid"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var Ser *Server

type Server struct {
	Ip        string
	IpVersion string
	Port      string
	ServerID  string
	upgrader  websocket.Upgrader
	server    *http.Server
	wg        sync.WaitGroup
}

func NewServer(ip, port string) *Server {
	serverID := uid.GetStringUUID("")
	return &Server{
		Ip:       ip,
		Port:     port,
		ServerID: serverID,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		wg: sync.WaitGroup{},
		server: &http.Server{
			Addr:         ":8080",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func Start(wg *sync.WaitGroup) {

	Ser = NewServer("errand.kuaixg.cn", "8080")
	r := RegistHttpRouter()
	Ser.server.Handler = r

	router.RegistRpcHandler_v2()
	fmt.Println("Server Start")
	err := Ser.server.ListenAndServe()
	if err != nil {
		wg.Done()
		fmt.Println("Server Error ", err)
		return
	}
	wg.Done()
	fmt.Println("Server Exit ")
}

func (s *Server) Close() {
	conn.CloseConnManager()
	s.server.Close()
}

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
