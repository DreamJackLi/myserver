package router

import (
	"robot-server/api/base"
	"robot-server/internal/core/message"
)

type FuncHandler func(*message.Message) *message.Message
type FuncHandler_v2 func(*base.Msg) *base.Msg

var apiMap map[int32]FuncHandler

var apiMap_v2 map[int32]FuncHandler_v2

func init() {
	apiMap = make(map[int32]FuncHandler)
	apiMap_v2 = make(map[int32]FuncHandler_v2)
}

func GetHandler(api int32) FuncHandler {

	f, ok := apiMap[api]

	if ok {
		return f
	} else {
		return nil
	}
}

func RegistRpcHandler() {
	// apiMap[int32(base.Api_GetEmpInfo)] = func(m *message.Message) *message.Message {
	//   fmt.Println("RegistRpcHandler Api_GetEmpInfo")
	//   // grpc 调用
	//   in := &apiemploy.GetInfoReq{}
	//   err := in.Unmarshal(m.Body)
	//   if err != nil {
	//     fmt.Println("Unmarshal Error")
	//     return nil
	//   }
	//   fmt.Println(in.ID)
	//   // in.Head = m.Head
	//   // grpc 调用
	//   cli := regist.GetCliConn(regist.ServerName[regist.EMPLOY])
	//   if cli == nil {
	//     fmt.Println("Employ Cli is nil")
	//     return nil
	//   }
	//   ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//   defer cancel()
	//   out := &apiemploy.GetInfoRes{}
	//   err = cli.Invoke(ctx, "/apiemploy.Employ/GetInfo", in, out)
	//   resB, _ := out.Marshal()
	//   resM := message.NewMessage(m.Head, resB)
	//   return resM
	// }

}

func GetHandler_v2(api int32) FuncHandler_v2 {

	f, ok := apiMap_v2[api]

	if ok {
		return f
	} else {
		return nil
	}
}

func RegistRpcHandler_v2() {
	// apiMap_v2[int32(base.Api_GetEmpInfo)] = getEmployInfo
	apiMap_v2[int32(base.Api_WebLogin)] = webLogin
	apiMap_v2[int32(base.Api_AddEmployeeInfo)] = addEmployeeInfo
	apiMap_v2[int32(base.Api_UpdateEmployeeInfo)] = updateEmployeeInfo
	apiMap_v2[int32(base.Api_GetEmployeeInfo)] = getEmployeeInfo
	apiMap_v2[int32(base.Api_AddEmployeeHealthInfo)] = addEmployeeHealthInfo
	apiMap_v2[int32(base.Api_UpdateEmployeeHealthInfo)] = updateEmployeeHealthInfo
	apiMap_v2[int32(base.Api_GetEmployeeHealthInfo)] = getEmployeeHealthInfo
	apiMap_v2[int32(base.Api_AddEmployeeHealthRecord)] = addEmployeeHealthRecord

	apiMap_v2[int32(base.Api_AddCompanyInfo)] = addCompanyInfo
	apiMap_v2[int32(base.Api_UpdateCompanyInfo)] = updateCompanyInfo
	apiMap_v2[int32(base.Api_GetCompanyInfo)] = getCompanyInfo
	apiMap_v2[int32(base.Api_AddPreventionRecord)] = addPreventionRecord
	apiMap_v2[int32(base.Api_UpdatePreventionRecord)] = updatePreventionRecord
	apiMap_v2[int32(base.Api_GetPreventionRecord)] = getPreventionRecord

	apiMap_v2[int32(base.Api_GetQiniuToken)] = getQiniuToken

	apiMap_v2[int32(base.Api_GetProvinces)] = getProvinces
	apiMap_v2[int32(base.Api_GetCity)] = getCity
	apiMap_v2[int32(base.Api_GetDistrict)] = getDistrict
	apiMap_v2[int32(base.Api_GetStreet)] = getStreet
	apiMap_v2[int32(base.Api_GetCommunity)] = getCommunity

	apiMap_v2[int32(base.Api_QueryEmployInfoByUser)] = queryEmployInfoByUser

	apiMap_v2[int32(base.Api_GetCompanyInfoList)] = getCompanyInfoList
	apiMap_v2[int32(base.Api_DeleteCompanyInfo)] = deleteCompanyInfo
	apiMap_v2[int32(base.Api_GetPreventionRecordList)] = getPreventionRecordList

	apiMap_v2[int32(base.Api_GetEmployeeHealthRecordList)] = getEmployeeHealthRecordList
	apiMap_v2[int32(base.Api_DeleteEmployInfo)] = deleteEmployInfo
	apiMap_v2[int32(base.Api_GetEmployeeHealthRecordListByCompanyId)] = getEmployeeHealthRecordListByCompanyId

	apiMap_v2[int32(base.Api_Export)] = export
}
