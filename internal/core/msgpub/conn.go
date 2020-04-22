package msgpub

import (
	"time"

	nats "github.com/nats-io/nats.go"
)

func Conn() {
	conn, err := nats.Connect(nats.DefaultOptions.Url)
	if err != nil {

	}

	msg, err := conn.Request("Test", []byte("Helloworld"), 3*time.Second)

	if err != nil {

	}
	msg.Respond([]byte("Helloworld"))

}
