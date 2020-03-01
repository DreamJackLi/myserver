package router

import (
	"context"
	"fmt"
	"myserver/api/apicompany"
	"myserver/api/base"
	"myserver/internal/core/message"
	"myserver/internal/core/regist"
	"time"
)

func addCompanyInfo(m *base.Msg) *base.Msg {
	in := &apicompany.AddCompanyInfoReq{}
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
	out := &apicompany.AddCompanyInfoRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/AddCompanyInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM

}

func updateCompanyInfo(m *base.Msg) *base.Msg {
	in := &apicompany.UpdateCompanyInfoReq{}
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
	out := &apicompany.UpdateCompanyInfoRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/UpdateCompanyInfoInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM

}

func getCompanyInfo(m *base.Msg) *base.Msg {
	in := &apicompany.GetCompanyInfoReq{}
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
	out := &apicompany.GetCompanyInfoRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/GetCompanyInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM

}

func addPreventionRecord(m *base.Msg) *base.Msg {
	in := &apicompany.AddPreventionRecordReq{}
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
	out := &apicompany.AddPreventionRecordRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/AddPreventionRecord", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func updatePreventionRecord(m *base.Msg) *base.Msg {
	in := &apicompany.UpdatePreventionRecordReq{}
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
	out := &apicompany.UpdatePreventionRecordRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/UpdatePreventionRecord", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getPreventionRecord(m *base.Msg) *base.Msg {
	in := &apicompany.GetPreventionRecordReq{}
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
	out := &apicompany.GetPreventionRecordRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/GetPreventionRecord", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getCompanyInfoList(m *base.Msg) *base.Msg {
	in := &apicompany.GetCompanyInfoListReq{}
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
	out := &apicompany.GetCompanyInfoListRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/GetCompanyInfoList", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func deleteCompanyInfo(m *base.Msg) *base.Msg {
	in := &apicompany.DeleteCompanyInfoReq{}
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
	out := &apicompany.DeleteCompanyInfoRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/DeleteCompanyInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getPreventionRecordList(m *base.Msg) *base.Msg {
	in := &apicompany.GetPreventionRecordListReq{}
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
	out := &apicompany.GetPreventionRecordListRes{}
	err = cli.Invoke(ctx, "/apicompany.Company/GetPreventionRecordList", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}
