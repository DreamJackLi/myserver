package message

import "robot-server/api/base"

const (
	HEADLEN = 4
	BODYLEN = 8
)

type Message struct {
	// 这里后面加个 MD5 校验啥的
	// HeadLen int32
	// BodyLen int64
	Head *base.Head
	Body []byte
}

func NewMessage(head *base.Head, body []byte) *Message {
	return &Message{
		Head: head,
		Body: body,
	}
}

func NewMessage_v2(apiID base.Api, data []byte) *base.Msg {
	return &base.Msg{
		ApiId: apiID,
		Data:  data,
	}
}

func BuildMsg(b []byte) (*base.Msg, error) {
	m := &base.Msg{}

	err := m.Unmarshal(b)

	if err != nil {
		return nil, err
	} else {
		return m, nil
	}

}
