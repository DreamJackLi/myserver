package mypaladin

import (
	"flag"
	"fmt"
)

var (
	DefaultClient Client
	confPath      string
)

func init() {
	flag.StringVar(&confPath, "../conf/", "", "default config path")
}

func Init(arg ...interface{}) (err error) {
	if confPath != "" {
		DefaultClient, err = NewFile(confPath)
	} else {
		return fmt.Errorf("please enter config file path")
	}
	return nil
}
