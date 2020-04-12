package log

import (
	"io"
	"os"
)

type Config struct {
	Family string
	Host   string

	Stdout bool

	Dir            string
	FileBufferSize int64
	MaxLogFile     int
	RotateSize     int64

	Module map[string]int32
	Filter []string
}

type Render interface {
	Render(io.Writer, map[string]interface{}) error
	RenderString(map[string]interface{}) string
}

var (
	h Handler
	c *Config
)

func init() {
	host, _ := os.Hostname()

	c = &Config{
		Host: host,
	}

	h = newHandler([]string{})
}

// func errIncr(lv Level , source string) {
//   if lv == _errorLevel {

//   }
// }
