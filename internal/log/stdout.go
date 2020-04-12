package log

import (
	"context"
	"os"
	"time"
)

const defaultPattern = "%L %d-%T %f %M"

type StdoutHandler struct {
	render Render
}

func NewStdout() *StdoutHandler {
	return &StdoutHandler{render: newPatternRender(defaultPattern)}
}

func (h *StdoutHandler) Log(ctx context.Context, lv Level, args ...D) {
	d := toMap(args...)

	// addExtraField(ctx , d)

	d[_time] = time.Now().Format(_timeFormat)
	h.render.Render(os.Stdout, d)

	os.Stderr.Write([]byte("\n"))
}
