package sig

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type CancelFunc func(os.Signal)

func Signal(cancelFunc CancelFunc) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	s := <-c

	cancelFunc(s)

	fmt.Println("Signal Close ", s.String())
}
