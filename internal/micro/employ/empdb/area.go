package empdb

import (
	"robot-server/api/apiarea"
	"robot-server/internal/core/coredb"
)

func GetProvinces(in *apiarea.Req) []*apiarea.Province {
	provinces := make([]*apiarea.Province, 0)
	coredb.Get().Table("t_province").Where(" country_id = ?", in.CountryId).Find(&provinces)
	return provinces
}

func GetCity(in *apiarea.Req) []*apiarea.City {
	cites := make([]*apiarea.City, 0)
	coredb.Get().Table("t_city").Where(" province_id = ?", in.ProvinceId).Find(&cites)
	return cites
}

func GetDistrict(in *apiarea.Req) []*apiarea.District {
	districts := make([]*apiarea.District, 0)
	coredb.Get().Table("t_area").Where(" city_id = ?", in.CityId).Find(&districts)
	return districts
}

func GetStreet(in *apiarea.Req) []*apiarea.Street {
	streets := make([]*apiarea.Street, 0)
	coredb.Get().Table("streets").Where(" district_id = ?", in.DistrictId).Find(&streets)
	return streets
}

func GetCommunity(in *apiarea.Req) []*apiarea.Community {
	communities := make([]*apiarea.Community, 0)
	coredb.Get().Table("communities").Where(" street_id = ?", in.StreetId).Find(&communities)
	return communities
}
