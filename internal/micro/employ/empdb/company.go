package empdb

import (
	"bytes"
	"fmt"
	"robot-server/api/apicompany"
	"robot-server/internal/core/coredb"
	"robot-server/tool"
	"strings"
	"time"
)

func AddCompanyInfo(company *apicompany.CompanyInfo) error {
	company.IsDelete = 1
	t := time.Now().Unix()
	company.CreateTime = t
	company.UpdateTime = t
	db := coredb.Get().Model(apicompany.CompanyInfo{}).Create(company)
	return db.Error
}

func UpdateCompanyInfo(company *apicompany.CompanyInfo) error {
	db := coredb.Get().Model(apicompany.CompanyInfo{}).Where("id = ?", company.Id).Update(company)
	return db.Error
}

func GetCompanyInfo(id int64) *apicompany.CompanyInfo {
	company := new(apicompany.CompanyInfo)
	coredb.Get().Model(apicompany.CompanyInfo{}).Where("id = ?", id).Find(company)
	return company
}
func GetCompanyDetailAddress(companyIDs string) ([]*apicompany.CompanyInfo, error) {

	companys := make([]*apicompany.CompanyInfo, 0)
	db := coredb.Get().Raw(fmt.Sprintf("select id , detailed_address from company_infos where id in(%s)", companyIDs)).Find(&companys)

	if db.Error != nil {
		return nil, db.Error
	}
	return companys, nil
}

func AddPreventionRecordReq(record *apicompany.PreventionRecord) error {
	db := coredb.Get().Model(apicompany.PreventionRecord{}).Create(record)
	return db.Error
}

func UpdatePreventionRecord(record *apicompany.PreventionRecord) error {
	db := coredb.Get().Model(apicompany.PreventionRecord{}).Where("id = ?", record.Id).Update(record)
	return db.Error
}
func GetPreventionRecord(companyId int64) *apicompany.PreventionRecord {
	record := new(apicompany.PreventionRecord)
	coredb.Get().Model(apicompany.PreventionRecord{}).Where("company_id = ?", companyId).Last(record)
	return record
}

func CountGetCompanyInfoList(in *apicompany.GetCompanyInfoListReq) int32 {
	var total int32
	sql := " SELECT count(*) FROM company_infos  WHERE "
	buf := new(bytes.Buffer)
	buf.WriteString(" is_delete = 1 AND adcode = ? ")
	params := make([]interface{}, 0)
	params = append(params, strings.TrimSpace(in.Adcode))
	if in.StreetId != 0 {
		buf.WriteString(" AND street_id = ? ")
		params = append(params, in.StreetId)
	}
	if in.CommunityId != 0 {
		buf.WriteString(" AND community_id = ? ")
		params = append(params, in.CommunityId)
	}
	if tool.IsNotBlank(in.Query) {
		buf.WriteString(" AND (name LIKE ? OR phone LIKE ?) ")
		params = append(params, "%"+in.Query+"%", "%"+in.Query+"%")
	}
	coredb.Get().Raw(fmt.Sprintf("%s%s", sql, buf.String()), params...).Count(&total)
	return total
}

func GetCompanyInfoList(in *apicompany.GetCompanyInfoListReq) []*apicompany.CompanyInfoListVo {
	sql := " SELECT * FROM company_infos  LEFT JOIN (SELECT p1.*FROM prevention_records p1 LEFT JOIN prevention_records p2 ON (p1.company_id = p2.company_id AND p1.id < p2.id) WHERE p2.id IS NULL) p ON company_infos.id = p.company_id WHERE "
	companyInfoListVos := make([]*apicompany.CompanyInfoListVo, 0)
	buf := new(bytes.Buffer)
	buf.WriteString(" is_delete = 1 AND adcode = ? ")
	params := make([]interface{}, 0)
	params = append(params, strings.TrimSpace(in.Adcode))
	if in.StreetId != 0 {
		buf.WriteString(" AND street_id = ? ")
		params = append(params, in.StreetId)
	}
	if in.CommunityId != 0 {
		buf.WriteString(" AND community_id = ? ")
		params = append(params, in.CommunityId)
	}
	if tool.IsNotBlank(in.Query) {
		buf.WriteString(" AND (name LIKE ? OR phone LIKE ?) ")
		params = append(params, "%"+in.Query+"%", "%"+in.Query+"%")
	}
	coredb.Get().Raw(fmt.Sprintf("%s%s", sql, buf.String()), params...).Limit(in.Limit).Offset(in.Limit * (in.Page - 1)).Scan(&companyInfoListVos)
	return companyInfoListVos
}

func DeleteCompanyInfo(ids string) error {
	db := coredb.Get().Model(apicompany.CompanyInfo{}).Where(fmt.Sprintf("id in (%s)", ids)).Update(&apicompany.CompanyInfo{IsDelete: 2, UpdateTime: time.Now().Unix()})
	return db.Error
}

func CountGetPreventionRecordList(companyId int64) int32 {
	var total int32
	coredb.Get().Model(apicompany.PreventionRecord{}).Where("company_id = ?", companyId).Count(&total)
	return total
}

func GetPreventionRecordList(in *apicompany.GetPreventionRecordListReq) []*apicompany.PreventionRecord {
	preventionRecords := make([]*apicompany.PreventionRecord, 0)
	// coredb.Get().Model(apicompany.PreventionRecord{}).Where("company_id = ?", in.CompanyId).Limit(in.Limit).Offset(in.Limit * (in.Limit - 1)).Find(&preventionRecords)
	coredb.Get().Model(apicompany.PreventionRecord{}).Where("company_id = ?", in.CompanyId).Find(&preventionRecords)
	return preventionRecords
}
