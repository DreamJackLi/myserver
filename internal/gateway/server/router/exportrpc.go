package router

import (
	"context"
	"fmt"
	"robot-server/api/apiexport"
	"robot-server/api/base"
	"robot-server/internal/core/message"
	"robot-server/internal/core/regist"
	"time"
)

func export(m *base.Msg) *base.Msg {
	in := &apiexport.ExportReq{}
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
	out := &apiexport.ExportRes{}
	err = cli.Invoke(ctx, "/apiexport.Export/Export", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}
