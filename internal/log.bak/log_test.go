package log

import "testing"

func TestLog(t *testing.T) {
	field := []D{
		KVString("remote", "0.0.0.0"),
		KVString("remote111", "2222"),
		KVString("remote2222", "33333"),
		KVString("remote3333", "444444"),
		KVString("remote55555", "6555555"),
	}

	Info(nil, field...)

}
