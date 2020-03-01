package router

import (
	"context"
	"fmt"
	"myserver/api/apiarea"
	"myserver/api/base"
	"myserver/internal/core/message"
	"myserver/internal/core/regist"
	"time"
)

func getProvinces(m *base.Msg) *base.Msg {
	in := &apiarea.Req{}
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
	out := &apiarea.Resp{}
	err = cli.Invoke(ctx, "/apiarea.Area/GetProvinces", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM

}

func getCity(m *base.Msg) *base.Msg {
	in := &apiarea.Req{}
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
	out := &apiarea.Resp{}
	err = cli.Invoke(ctx, "/apiarea.Area/GetCity", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getDistrict(m *base.Msg) *base.Msg {
	in := &apiarea.Req{}
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
	out := &apiarea.Resp{}
	err = cli.Invoke(ctx, "/apiarea.Area/GetDistrict", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getStreet(m *base.Msg) *base.Msg {
	in := &apiarea.Req{}
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
	out := &apiarea.Resp{}
	err = cli.Invoke(ctx, "/apiarea.Area/GetStreet", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getCommunity(m *base.Msg) *base.Msg {
	in := &apiarea.Req{}
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
	out := &apiarea.Resp{}
	err = cli.Invoke(ctx, "/apiarea.Area/GetCommunity", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}
