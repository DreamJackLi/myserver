package message

import (
	"encoding/binary"
	"errors"
	"fmt"
	"robot-server/api/base"
)

const (
	MAX_HEAD_LEN = 65535
)

func UnPacker(b []byte) (*Message, error) {
	if len(b) < HEADLEN+BODYLEN {
		return nil, errors.New("Data too small")
	}
	headLen := binary.BigEndian.Uint32(b[0:HEADLEN])
	if headLen > MAX_HEAD_LEN {
		return nil, errors.New("Head too big")
	}
	// bodyLen := binary.BigEndian.Uint64(b[HEADLEN:BODYLEN])
	h := &base.Head{}

	err := h.Unmarshal(b[HEADLEN+BODYLEN : HEADLEN+BODYLEN+headLen])
	if err != nil {
		return nil, err
	}

	return NewMessage(h, b[HEADLEN+BODYLEN+headLen:]), nil
}

func Pack(m *Message) ([]byte, error) {
	if m == nil || m.Head == nil {
		fmt.Println("Pack Error Message illegal ")
		return nil, errors.New("Message illegal")
	}
	hM, err := m.Head.Marshal()
	if err != nil {
		return nil, err
	}
	fmt.Println(hM)
	hLen := len(hM)
	bLen := len(m.Body)
	b := make([]byte, HEADLEN+BODYLEN, HEADLEN+BODYLEN+hLen+bLen)
	binary.BigEndian.PutUint32(b[0:HEADLEN], uint32(hLen))

	binary.BigEndian.PutUint64(b[HEADLEN:HEADLEN+BODYLEN], uint64(bLen))

	b = append(b, hM...)
	b = append(b, m.Body...)

	return b, nil
}

func UnPacker_v2(b []byte) (*base.Msg, error) {
	fmt.Println("UnPacker data is ", len(b))
	m := &base.Msg{}
	err := m.Unmarshal(b)

	if err != nil {
		return nil, err
	}
	return m, nil
}

func Pack_v2(m *base.Msg) ([]byte, error) {

	if m == nil || m.ApiId == 0 || m.Data == nil || len(m.Data) == 0 {
		return nil, errors.New("Msg Format illegal")
	}

	b, err := m.Marshal()

	if err != nil {
		return nil, err
	}
	return b, nil
}
