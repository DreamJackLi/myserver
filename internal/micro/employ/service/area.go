package service

import (
	"context"
	"robot-server/api/apiarea"
	"robot-server/api/base"
	"robot-server/internal/core/errorcode"
	"robot-server/internal/micro/employ/empdb"
)

type AreaServer struct{}

// 获取省
func (a *AreaServer) GetProvinces(ctx context.Context, in *apiarea.Req) (*apiarea.Resp, error) {
	// 参数验证
	if in == nil {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CountryId == 0 {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	provinces := empdb.GetProvinces(in)
	return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Provinces: provinces}, nil
}

// 获取市
func (a *AreaServer) GetCity(ctx context.Context, in *apiarea.Req) (*apiarea.Resp, error) {
	if in == nil {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.ProvinceId == 0 {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	cities := empdb.GetCity(in)
	return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Cites: cities}, nil
}

// 获取县/区
func (a *AreaServer) GetDistrict(ctx context.Context, in *apiarea.Req) (*apiarea.Resp, error) {
	if in == nil {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.CityId == 0 {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	districts := empdb.GetDistrict(in)
	return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Districts: districts}, nil
}

// 获取街道/镇
func (a *AreaServer) GetStreet(ctx context.Context, in *apiarea.Req) (*apiarea.Resp, error) {
	if in == nil {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.DistrictId == 0 {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	streets := empdb.GetStreet(in)
	return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Streets: streets}, nil
}

// 获取社区
func (a *AreaServer) GetCommunity(ctx context.Context, in *apiarea.Req) (*apiarea.Resp, error) {
	if in == nil {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if in.StreetId == 0 {
		return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	communities := empdb.GetCommunity(in)
	return &apiarea.Resp{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}, Communities: communities}, nil
}
