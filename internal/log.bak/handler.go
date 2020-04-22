package log

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

const (
	_timeFormat = "2006-01-02T15:04:05.999999"

	_levelValue = "level_value"
	_source     = "source"
	_time       = "time"
	_level      = "level"
)

type Handler interface {
	Log(context.Context, Level, ...D)
	SetFormat(string)
	Close() error
}

func newHandler(filters []string, handler ...Handler) *Handlers {
	set := make(map[string]struct{})

	for _, k := range filters {
		set[k] = struct{}{}
	}
	return &Handlers{filters: set, handlers: handler}
}

type Handlers struct {
	filters  map[string]struct{}
	handlers []Handler
}

func (hs Handlers) Log(ctx context.Context, lv Level, d ...D) {
	// hasSource := false

	for i := range d {
		if _, ok := hs.filters[d[i].Key]; ok {
			d[i].Value = "***"
		}

		// if d[i].Key == _source {
		//   hasSource = true
		// }
	}

	// if !hasSource {
	//   fn := funcName(3)
	//   errIncr()
	// }

	d = append(d, KV(_time, time.Now()), KVInt64(_levelValue, int64(lv)), KVString(_level, lv.String()))
	for _, h := range hs.handlers {
		h.Log(ctx, lv, d...)
	}
}

func (hs Handlers) Close() (err error) {
	for _, h := range hs.handlers {
		if e := h.Close(); e != nil {
			err = errors.WithStack(e)
		}
	}
	return
}

func (hs Handlers) SetFormat(format string) {
	for _, h := range hs.handlers {
		h.SetFormat(format)
	}
}
