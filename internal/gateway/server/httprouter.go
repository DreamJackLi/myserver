package server

import (
	"fmt"

	"github.com/gorilla/mux"
)

type RouterHandler func()

func GetHandler(apiID int64) RouterHandler {
	return func() {}
}

func RegistHttpRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/begin", buildConn)
	fmt.Println("HandleFunc /begin")

	router.HandleFunc("/", Ser.hello)
	return router
}
