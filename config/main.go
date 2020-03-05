package main

import (
	"flag"
	"fmt"
	"myserver/config/mypaladin"
	"time"
)

func main() {
	flag.Parse()

	err := mypaladin.Init()

	if err != nil {
		fmt.Println(err)
	}
	t := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-t.C:
			s, _ := mypaladin.DefaultClient.Get("config_prev.yaml").String()
			fmt.Println(s)
		}
	}
}
