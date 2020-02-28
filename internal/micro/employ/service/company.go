package service

import (
	"context"
	"fmt"
	"robot-server/api/apicompany"
	"robot-server/api/base"
	"robot-server/internal/core/errorcode"
	"robot-server/internal/micro/employ/empdb"
	"robot-server/tool"
	"time"
)

type CompanyServer struct{}

// 添加企业基本信息
func (c *CompanyServer) AddCompanyInfo(ctx context.Context, in *apicompany.AddCompanyInfoReq) (*apicompany.AddCompanyInfoRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Name) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Phone) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Province) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ProvinceId == 0 {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.City) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CityId == 0 {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.District) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.DistrictId == 0 {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Street) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.StreetId == 0 {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Community) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CommunityId == 0 {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DetailedAddress) {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 保存到数据库
	company := new(apicompany.CompanyInfo)
	tool.CopyProperties(in, company)
	err := empdb.AddCompanyInfo(company)
	if err != nil {
		return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.InsertError}}, nil
	}
	return &apicompany.AddCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Id: company.Id}, nil
}

// 编辑企业基本信息
func (c *CompanyServer) UpdateCompanyInfo(ctx context.Context, in *apicompany.UpdateCompanyInfoReq) (*apicompany.UpdateCompanyInfoRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Name) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Id == 0 {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Phone) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Province) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ProvinceId == 0 {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.City) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CityId == 0 {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.District) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.DistrictId == 0 {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Street) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.StreetId == 0 {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Community) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CommunityId == 0 {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DetailedAddress) {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 到数据库修改信息
	company := new(apicompany.CompanyInfo)
	tool.CopyProperties(in, company)
	err := empdb.UpdateCompanyInfo(company)
	if err != nil {
		return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.UpdateError}}, nil
	}
	return &apicompany.UpdateCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 获取公司基本信息
func (c *CompanyServer) GetCompanyInfo(ctx context.Context, in *apicompany.GetCompanyInfoReq) (*apicompany.GetCompanyInfoRes, error) {
	if in == nil {
		return &apicompany.GetCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	fmt.Println("GetCompanyInfo in ", in)
	if in.Id == 0 {
		return &apicompany.GetCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	companyInfo := empdb.GetCompanyInfo(in.Id)
	return &apicompany.GetCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, CompanyInfo: companyInfo}, nil
}

// 添加企业防控信息
func (c *CompanyServer) AddPreventionRecord(ctx context.Context, in *apicompany.AddPreventionRecordReq) (*apicompany.AddPreventionRecordRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ReturneeNum < 0 {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.MaterialImage) {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DisinfectionImage) {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.DisinfectionTime == 0 {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DisinfectionRange) {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DisinfectionOperator) {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 当天已有信息，不再做添加
	preventionRecord := empdb.GetPreventionRecord(in.CompanyId)
	if preventionRecord.Id != 0 && today(preventionRecord.DisinfectionTime) {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.Repeat}}, nil
	}
	// 到数据库保存信息
	record := new(apicompany.PreventionRecord)
	tool.CopyProperties(in, record)
	err := empdb.AddPreventionRecordReq(record)
	if err != nil {
		return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.InsertError}}, nil
	}
	return &apicompany.AddPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Id: record.Id}, nil
}

func today(t int64) bool {
	format := "2006-01-02"
	if time.Unix(t, 0).Format(format) == time.Now().Format(format) {
		return true
	}
	return false
}

// 编辑企业防控信息
func (c *CompanyServer) UpdatePreventionRecord(ctx context.Context, in *apicompany.UpdatePreventionRecordReq) (*apicompany.UpdatePreventionRecordRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Id == 0 {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ReturneeNum < 0 {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.MaterialImage) {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DisinfectionImage) {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.DisinfectionTime == 0 {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DisinfectionRange) {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.DisinfectionOperator) {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 到数据库保存信息
	record := new(apicompany.PreventionRecord)
	tool.CopyProperties(in, record)
	err := empdb.UpdatePreventionRecord(record)
	if err != nil {
		return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.UpdateError}}, nil
	}
	return &apicompany.UpdatePreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 获取企业防控信息
func (c *CompanyServer) GetPreventionRecord(ctx context.Context, in *apicompany.GetPreventionRecordReq) (*apicompany.GetPreventionRecordRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.GetPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apicompany.GetPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 获取数据库信息
	record := empdb.GetPreventionRecord(in.CompanyId)
	return &apicompany.GetPreventionRecordRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, PreventionRecord: record}, nil
}

// 获取企业信息列表
func (c *CompanyServer) GetCompanyInfoList(ctx context.Context, in *apicompany.GetCompanyInfoListReq) (*apicompany.GetCompanyInfoListRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.GetCompanyInfoListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Adcode) {
		return &apicompany.GetCompanyInfoListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Page == 0 {
		return &apicompany.GetCompanyInfoListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Limit == 0 {
		return &apicompany.GetCompanyInfoListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// 获取列表信息数量
	total := empdb.CountGetCompanyInfoList(in)
	if total <= 0 {
		return &apicompany.GetCompanyInfoListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: total}, nil
	}
	companyInfoList := empdb.GetCompanyInfoList(in)
	return &apicompany.GetCompanyInfoListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, CompanyInfoList: companyInfoList, Total: total}, nil
}

// 删除企业基本信息
func (c *CompanyServer) DeleteCompanyInfo(ctx context.Context, in *apicompany.DeleteCompanyInfoReq) (*apicompany.DeleteCompanyInfoRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.DeleteCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.Ids == "" {
		return &apicompany.DeleteCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}

	err := empdb.DeleteCompanyInfo(in.Ids)
	if err != nil {
		return &apicompany.DeleteCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.DeleteError}}, nil
	}
	return &apicompany.DeleteCompanyInfoRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 获取企业防控信息列表
func (c *CompanyServer) GetPreventionRecordList(ctx context.Context, in *apicompany.GetPreventionRecordListReq) (*apicompany.GetPreventionRecordListRes, error) {
	// 数据验证
	if in == nil {
		return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CompanyId == 0 {
		return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	// if in.Page == 0 {
	//   return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	// }
	// if in.Limit == 0 {
	//   return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	// }
	// total := empdb.CountGetPreventionRecordList(in.CompanyId)
	// if total <= 0 {
	// return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Total: total}, nil
	// }
	preventionRecordList := empdb.GetPreventionRecordList(in)
	if len(preventionRecordList) == 0 {

		return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
	}
	preventionRecordListVos := make([]*apicompany.PreventionRecordListVo, len(preventionRecordList))
	tool.CopyProperties(&preventionRecordList, &preventionRecordListVos)
	return &apicompany.GetPreventionRecordListRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, PreventionRecordList: preventionRecordListVos}, nil
}

func buildComMap(c []*apicompany.CompanyInfo) map[int64]*apicompany.CompanyInfo {
	m := make(map[int64]*apicompany.CompanyInfo)

	for _, v := range c {
		temp := v
		m[v.Id] = temp
	}

	return m

}
