package service

import (
	"context"
	"fmt"
	"robot-server/api/apiemployee"
	"robot-server/api/base"
	"robot-server/config"
	"robot-server/internal/core/errorcode"
	"robot-server/internal/micro/employ/empdb"
	"robot-server/tool"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

type EmployServer struct{}

func (e *EmployServer) GetQiniuToken(ctx context.Context, in *apiemployee.GetQiniuTokenReq) (*apiemployee.GetQiniuTokenRes, error) {

	// 做一些校验
	//获取七牛云token
	bucket := config.Get().Qiniu.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 3600 * 24 * 356 * 2 //示例两年有效期
	mac := qbox.NewMac(config.Get().Qiniu.AccessKey, config.Get().Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	return &apiemployee.GetQiniuTokenRes{
		Base: &base.BaseRes{
			ErrorCode: errorcode.Succeed,
		},
		Token: upToken,
	}, nil

}

func (e *EmployServer) WebLogin(ctx context.Context, in *apiemployee.WebLoginReq) (*apiemployee.WebLoginRes, error) {
	if in.UserName == "" || in.Pwd == "" {
		fmt.Println(in.UserName)
		fmt.Println(in.Pwd)
		return &apiemployee.WebLoginRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	info, err := empdb.WebLogin(in.UserName)

	if err != nil {

		return &apiemployee.WebLoginRes{Base: &base.BaseRes{ErrorCode: errorcode.FindError}}, nil
	}

	if info.UserPwd != in.Pwd {
		return &apiemployee.WebLoginRes{Base: &base.BaseRes{ErrorCode: errorcode.PwdError}}, nil
	}

	return &apiemployee.WebLoginRes{
		Base: &base.BaseRes{
			ErrorCode: errorcode.Succeed,
		},
		Info: info,
	}, nil
}

// 添加员工基本信息
func (e *EmployServer) AddEmployeeInfo(ctx context.Context, in *apiemployee.AddEmployeeInfoReq) (*apiemployee.AddEmployeeInfoRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Name) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Phone) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyProvince) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyProvinceId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyCity) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyCityId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyDistrict) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyDistrictId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyStreet) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyStreetId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyCommunity) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyCommunityId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyName) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.IdCardNumber) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Gender == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Age < 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceProvince) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceProvinceId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceCity) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceCityId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceDistrict) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceDistrictId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceStreet) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceStreetId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceCommunity) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceCommunityId == 0 {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceDetailedAddress) {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 数据库保存数据
	employeeInfo := new(apiemployee.EmployeeInfo)
	tool.CopyProperties(in, employeeInfo)
	err := empdb.AddEmployeeInfo(employeeInfo)
	if err != nil {
		return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.InsertError}}, nil
	}
	return &apiemployee.AddEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 编辑员工基本信息
func (e *EmployServer) UpdateEmployeeInfo(ctx context.Context, in *apiemployee.UpdateEmployeeInfoReq) (*apiemployee.UpdateEmployeeInfoRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Id == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Name) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Phone) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyProvince) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyProvinceId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyCity) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyCityId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyDistrict) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyDistrictId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyStreet) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyStreetId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyCommunity) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyCommunityId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.CompanyName) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.IdCardNumber) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Gender == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Age < 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceProvince) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceProvinceId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceCity) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceCityId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceDistrict) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceDistrictId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceStreet) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceStreetId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceCommunity) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ResidenceCommunityId == 0 {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ResidenceDetailedAddress) {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 数据库保存数据
	employeeInfo := new(apiemployee.EmployeeInfo)
	tool.CopyProperties(in, employeeInfo)
	err := empdb.UpdateEmployeeInfo(employeeInfo)
	if err != nil {
		return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.UpdateError}}, nil
	}
	return &apiemployee.UpdateEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 获取员工基本信息
func (e *EmployServer) GetEmployeeInfo(ctx context.Context, in *apiemployee.GetEmployeeInfoReq) (*apiemployee.GetEmployeeInfoRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.GetEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Id == 0 {
		return &apiemployee.GetEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 数据库获取数据
	employeeInfo := empdb.GetEmployeeInfo(in.Id)
	return &apiemployee.GetEmployeeInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, EmployeeInfo: employeeInfo}, nil
}

// 添加员工健康信息
func (e *EmployServer) AddEmployeeHealthInfo(ctx context.Context, in *apiemployee.AddEmployeeHealthInfoReq) (*apiemployee.AddEmployeeHealthInfoRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.EmployeeId == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.SymptomWithin14Days) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.FormHuBei == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Touch_NCP == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.BeenToHuBei == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ReturnChengDuTime == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ReturnChengDuDetailedAddress) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ReturnChengDuTransport) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.TransportDetailInfo) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.TouchHuBeiPeopleWithin14Days == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.FromChengDuOutside == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ThisPlaceTime == 0 {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ReturnThisPlaceTransport) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Profession) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.LifeTrace) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.SignImage) {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 保存数据
	employeeHealthInfo := new(apiemployee.EmployeeHealthInfo)
	tool.CopyProperties(in, employeeHealthInfo)
	err := empdb.AddEmployeeHealthInfo(employeeHealthInfo)
	if err != nil {
		return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.InsertError}}, nil
	}
	return &apiemployee.AddEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 修改员工健康信息
func (e *EmployServer) UpdateEmployeeHealthInfo(ctx context.Context, in *apiemployee.UpdateEmployeeHealthInfoReq) (*apiemployee.UpdateEmployeeHealthInfoRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.EmployeeId == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.SymptomWithin14Days) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.FormHuBei == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Touch_NCP == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.BeenToHuBei == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ReturnChengDuTime == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ReturnChengDuDetailedAddress) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ReturnChengDuTransport) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.TransportDetailInfo) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.TouchHuBeiPeopleWithin14Days == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.FromChengDuOutside == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ThisPlaceTime == 0 {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.ReturnThisPlaceTransport) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Profession) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.LifeTrace) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.SignImage) {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 保存数据
	employeeHealthInfo := new(apiemployee.EmployeeHealthInfo)
	tool.CopyProperties(in, employeeHealthInfo)
	err := empdb.UpdateEmployeeHealthInfo(employeeHealthInfo)
	if err != nil {
		return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.UpdateError}}, nil
	}
	return &apiemployee.UpdateEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 获取员工健康信息
func (e *EmployServer) GetEmployeeHealthInfo(ctx context.Context, in *apiemployee.GetEmployeeHealthInfoReq) (*apiemployee.GetEmployeeHealthInfoRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.GetEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.EmployeeId == 0 {
		return &apiemployee.GetEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	employeeHealthInfo := empdb.GetEmployeeHealthInfo(in.EmployeeId)
	return &apiemployee.GetEmployeeHealthInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, EmployeeHealthInfo: employeeHealthInfo}, nil
}

// 添加/更新员工健康台账
func (e *EmployServer) AddEmployeeHealthRecord(ctx context.Context, in *apiemployee.AddEmployeeHealthRecordReq) (*apiemployee.AddEmployeeHealthRecordRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.EmployeeId == 0 {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Department) {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Temperature < 0 {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.EpidemicSymptom == 0 {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DetectionResult) {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DetectionOperator) {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	employeeHealthRecord := new(apiemployee.EmployeeHealthRecord)
	tool.CopyProperties(in, employeeHealthRecord)
	employeeHealthRecord.CreatedTime = time.Now().Unix()
	err := empdb.AddEmployeeHealthRecord(employeeHealthRecord)
	if err != nil {
		return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.InsertError}}, nil
	}
	return &apiemployee.AddEmployeeHealthRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

/// 根据公司id，获取员工最新健康台账
func (e *EmployServer) GetEmployeeHealthRecordListByCompanyId(ctx context.Context, in *apiemployee.GetEmployeeHealthRecordListByCompanyIdReq) (*apiemployee.GetEmployeeHealthRecordListByCompanyIdRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.GetEmployeeHealthRecordListByCompanyIdRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apiemployee.GetEmployeeHealthRecordListByCompanyIdRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	employeeHealthRecords := empdb.GetEmployeeHealthRecordListByCompanyId(in.CompanyId)
	return &apiemployee.GetEmployeeHealthRecordListByCompanyIdRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, EmployeeHealthRecords: employeeHealthRecords}, nil
}

// 获取员工健康台账历史记录
func (e *EmployServer) GetEmployeeHealthRecordList(ctx context.Context, in *apiemployee.GetEmployeeHealthRecordListReq) (*apiemployee.GetEmployeeHealthRecordListRes, error) {
	// 参数验证
	if in == nil {
		return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}

	fmt.Println(in)
	if in.EmployeeId == 0 {
		return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// if in.Page == 0 {
	//   return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	// }
	// if in.Limit == 0 {
	//   return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	// }
	// total := empdb.CountGetEmployeeHealthRecordList(in.EmployeeId)

	// if total <= 0 {
	//   return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: total}, nil
	// }
	employeeHealthRecords := empdb.GetEmployeeHealthRecordList(in)
	if len(employeeHealthRecords) == 0 {

		return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
	}
	healthRecordList := make([]*apiemployee.EmployeeHealthRecordListVo, len(employeeHealthRecords))
	tool.CopyProperties(&employeeHealthRecords, &healthRecordList)
	return &apiemployee.GetEmployeeHealthRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, EmployeeHealthRecordList: healthRecordList}, nil
}

func (e *EmployServer) QueryEmployInfoByUser(ctx context.Context, in *apiemployee.QueryEmployInfoByUserReq) (*apiemployee.QueryEmployInfoByUserRes, error) {

	if in == nil || in.CommunityId == 0 || in.StreetId == 0 || in.UserType == 0 {
		return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	fmt.Println("in QueryEmployInfoByUser  ", in)
	var empInfos []*apiemployee.EmployeeInfo
	var err error
	var total int64
	// 1 表示街道账号 2 表示社区账号  3表示超级账号
	if in.UserType == 1 || in.UserType == 3 {
		empInfos, err = empdb.GetStreetInfo(in.StreetId, in.LastRecordId, in.Page, in.Limit, in.Query)
		if err != nil {
			return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 1}, nil
		}
		total, err = empdb.GetStreetInfoTotal(in.StreetId, in.Query)

		if err != nil {
			return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 1}, nil
		}

	} else if in.UserType == 2 {

		empInfos, err = empdb.GetCommunityInfo(in.CommunityId, in.LastRecordId, in.Page, in.Limit, in.Query)
		if err != nil {
			return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 1}, nil
		}
		total, err = empdb.GetCommunityInfoTotal(in.CommunityId, in.Query)
		if err != nil {
			return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 1}, nil
		}
	}

	if len(empInfos) == 0 {
		return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 1}, nil
	}

	ids, companyIds := getIds(empInfos)

	empHealthInfos, err := empdb.GetEmployeeHealthInfos(ids)

	if err != nil {

		return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 1}, nil
	}

	comAddress, err := empdb.GetCompanyDetailAddress(companyIds)

	if err != nil {
		return &apiemployee.QueryEmployInfoByUserRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: 0}, nil
	}

	empHealthMap := buildEmpMap(empHealthInfos)

	comMap := buildComMap(comAddress)
	res := make([]*apiemployee.QueryEmployInfoByUserMode, 0, len(empInfos))

	for _, v := range empInfos {

		t := &apiemployee.QueryEmployInfoByUserMode{
			Name:                     v.Name,
			Phone:                    v.Phone,
			Age:                      v.Age,
			Gender:                   v.Gender,
			CompanyName:              v.CompanyName,
			CompanyId:                v.CompanyId,
			IdCardNumber:             v.IdCardNumber,
			CompanyCommunity:         v.CompanyCommunity,
			CompanyCommunityId:       v.CompanyCommunityId,
			ResidenceDetailedAddress: v.ResidenceDetailedAddress,
			ResidenceCommunity:       v.ResidenceCommunity,
			ResidenceCommunityId:     v.ResidenceCommunityId,
			CreatedTime:              v.CreateTime,
			UpdateTime:               v.UpdateTime,
			CompanyProvince:          v.CompanyProvince,
			CompanyProvinceId:        v.CompanyProvinceId,
			CompanyCity:              v.CompanyCity,
			CompanyCityId:            v.CompanyCityId,
			CompanyDistrict:          v.CompanyDistrict,
			CompanyDistrictId:        v.CompanyDistrictId,
		}
		h, ok := empHealthMap[v.Id]

		if ok {
			t.SymptomWithin14Days = h.SymptomWithin14Days
			t.FormHuBei = h.FormHuBei
			t.Touch_NCP = h.Touch_NCP
			t.BeenToHuBei = h.BeenToHuBei
			t.ReturnChengDuTime = h.ReturnChengDuTime
			t.ReturnChengDuDetailedAddress = h.ReturnChengDuDetailedAddress
			t.ReturnChengDuTransport = h.ReturnChengDuTransport
			t.TransportDetailInfo = h.TransportDetailInfo
			t.TouchHuBeiPeopleWithin14Days = h.TouchHuBeiPeopleWithin14Days
			t.LastTouchTime = h.LastTouchTime
			t.FromChengDuOutside = h.FromChengDuOutside
			t.ThisPlaceTime = h.ThisPlaceTime
			t.ReturnThisPlaceTransport = h.ReturnThisPlaceTransport
			t.Profession = h.Profession
			t.LifeTrace = h.LifeTrace
			t.SignImage = h.SignImage
			t.EmployeeId = v.Id
			t.ComeFromWenjiang = h.ComeFromWenjiang
			t.ComeFromAddress = h.ComeFromAddress
			t.ComeFromDetailedAddress = h.ComeFromDetailedAddress
		}

		c, ok := comMap[v.CompanyId]

		if ok {
			t.DetailedAddress = c.DetailedAddress
		}

		r, err := empdb.GetEmployeeHealthRecordLast(v.Id)

		if err == nil {
			t.Department = r.Department
			t.Temperature = r.Temperature
			t.EpidemicSymptom = r.EpidemicSymptom
			t.DetectionResult = r.DetectionResult
			t.DetectionOperator = r.DetectionOperator
		}

		res = append(res, t)
	}

	return &apiemployee.QueryEmployInfoByUserRes{
		Base: &base.BaseRes{
			ErrorCode: errorcode.Succeed,
		},
		Info:  res,
		Total: total,
	}, nil
}

func (e *EmployServer) DeleteEmployInfo(ctx context.Context, in *apiemployee.DeleteEmployInfoReq) (*apiemployee.DeleteEmployInfoRes, error) {
	if in == nil || in.EmpIds == "" {
		return &apiemployee.DeleteEmployInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}

	err := empdb.DeleteEmployInfo(in.EmpIds)

	if err != nil {
		return &apiemployee.DeleteEmployInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.FindError}}, nil
	}

	return &apiemployee.DeleteEmployInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

func getIds(empInfos []*apiemployee.EmployeeInfo) (string, string) {
	ids := ""
	companyIds := ""
	for i := 0; i < len(empInfos); i++ {
		ids += fmt.Sprintf("%d", empInfos[i].Id)
		companyIds += fmt.Sprintf("%d", empInfos[i].CompanyId)
		if i < len(empInfos)-1 {
			ids += ","
			companyIds += ","
		}
	}

	return ids, companyIds
}

func buildEmpMap(e []*apiemployee.EmployeeHealthInfo) map[int64]*apiemployee.EmployeeHealthInfo {
	m := make(map[int64]*apiemployee.EmployeeHealthInfo)

	for _, v := range e {
		temp := v
		m[v.EmployeeId] = temp
	}

	return m
}
