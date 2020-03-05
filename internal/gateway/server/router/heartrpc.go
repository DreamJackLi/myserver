package router

import (
	"context"
	"fmt"
	"myserver/api/apiemployee"
	"myserver/api/base"
	"myserver/internal/core/message"
	"myserver/internal/core/regist"
	"time"
)

func heartCheck(m *base.Msg) *base.Msg {
	in := &base.HeartReq{}
	err := in.Unmarshal(m.Data)
	if err != nil {
		fmt.Println("Unmarshal Error")
		return nil
	}
	// grpc 调用
	cli := regist.GetCliConn(regist.ServerName[regist.ALL])
	if cli == nil {
		fmt.Println("Employ Cli is nil")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	out := &apiemployee.WebLoginRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/WebLogin", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM

}
