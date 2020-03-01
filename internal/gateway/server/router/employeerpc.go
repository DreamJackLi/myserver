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

func webLogin(m *base.Msg) *base.Msg {
	in := &apiemployee.WebLoginReq{}
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
func addEmployeeInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.AddEmployeeInfoReq{}
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
	out := &apiemployee.AddEmployeeInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/AddEmployeeInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func updateEmployeeInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.UpdateEmployeeInfoReq{}
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
	out := &apiemployee.UpdateEmployeeInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/UpdateEmployeeInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getEmployeeInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.GetEmployeeInfoReq{}
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
	out := &apiemployee.GetEmployeeInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/GetEmployeeInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func addEmployeeHealthInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.AddEmployeeHealthInfoReq{}
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
	out := &apiemployee.AddEmployeeHealthInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/AddEmployeeHealthInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func updateEmployeeHealthInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.UpdateEmployeeHealthInfoReq{}
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
	out := &apiemployee.UpdateEmployeeHealthInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/UpdateEmployeeHealthInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getEmployeeHealthInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.GetEmployeeHealthInfoReq{}
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
	out := &apiemployee.GetEmployeeHealthInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/GetEmployeeHealthInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func addEmployeeHealthRecord(m *base.Msg) *base.Msg {
	in := &apiemployee.AddEmployeeHealthRecordReq{}
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
	out := &apiemployee.AddEmployeeHealthRecordRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/AddEmployeeHealthRecord", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getQiniuToken(m *base.Msg) *base.Msg {
	in := &apiemployee.GetQiniuTokenReq{}
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
	out := &apiemployee.GetQiniuTokenRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/GetQiniuToken", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func queryEmployInfoByUser(m *base.Msg) *base.Msg {
	in := &apiemployee.QueryEmployInfoByUserReq{}
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
	out := &apiemployee.QueryEmployInfoByUserRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/QueryEmployInfoByUser", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getEmployeeHealthRecordList(m *base.Msg) *base.Msg {
	in := &apiemployee.GetEmployeeHealthRecordListReq{}
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
	out := &apiemployee.GetEmployeeHealthRecordListRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/GetEmployeeHealthRecordList", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func deleteEmployInfo(m *base.Msg) *base.Msg {
	in := &apiemployee.DeleteEmployInfoReq{}
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
	out := &apiemployee.DeleteEmployInfoRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/DeleteEmployInfo", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}

func getEmployeeHealthRecordListByCompanyId(m *base.Msg) *base.Msg {
	in := &apiemployee.GetEmployeeHealthRecordListByCompanyIdReq{}
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
	out := &apiemployee.GetEmployeeHealthRecordListByCompanyIdRes{}
	err = cli.Invoke(ctx, "/apiemployee.Employee/GetEmployeeHealthRecordListByCompanyId", in, out)
	resB, _ := out.Marshal()
	resM := message.NewMessage_v2(m.ApiId, resB)
	return resM
}
