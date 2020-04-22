package log

import "myserver/internal/log/core"

type D core.Field

func KVString(key string, value string) D {
	return D{Key: key, Type: core.StringType, StringVal: value}
}

func KV(key string, value interface{}) D {
	return D{Key: key, Value: value}
}

func KVInt64(key string, value int64) D {
	return D{Key: key, Type: core.Int64Type, Value: value}
}
