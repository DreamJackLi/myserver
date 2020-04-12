package log

import (
	"math"
	"myserver/internal/log/core"
	"runtime"
	"strconv"
	"time"
)

func funcName(skip int) string {
	if _, file, lineNo, ok := runtime.Caller(skip); ok {
		return file + ":" + strconv.Itoa(lineNo)
	}
	return "unknown:0"
}

func toMap(args ...D) map[string]interface{} {
	d := make(map[string]interface{}, 10+len(args))

	for _, arg := range args {
		switch arg.Type {
		case core.Uint64Type, core.UintType:
			d[arg.Key] = arg.Int64Val
		case core.StringType:
			d[arg.Key] = arg.StringVal
		case core.Float32Type:
			d[arg.Key] = math.Float32frombits(uint32(arg.Int64Val))
		case core.Float64Type:
			d[arg.Key] = math.Float64frombits(uint64(arg.Int64Val))
		case core.DurationType:
			d[arg.Key] = time.Duration(arg.Int64Val)
		default:
			d[arg.Key] = arg.Value
		}
	}

	return d
}
