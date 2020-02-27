package empdb

import (
	"bytes"
	"errors"
	"fmt"
	"robot-server/api/apiemployee"
	"robot-server/internal/core/coredb"
	"robot-server/tool"

	"github.com/jinzhu/gorm"
)

func WebLogin(userName string) (*apiemployee.WebLoginInfo, error) {
	res := &apiemployee.WebLoginInfo{}
	db := coredb.Get().Raw("select * from login where user_name = ? ", userName).Find(res)
	if db.Error != nil {
		return nil, db.Error
	}

	return res, nil
}

func AddEmployeeInfo(employeeInfo *apiemployee.EmployeeInfo) error {
	db := coredb.Get().Model(apiemployee.EmployeeInfo{}).Create(employeeInfo)
	return db.Error
}

func UpdateEmployeeInfo(employeeInfo *apiemployee.EmployeeInfo) error {
	db := coredb.Get().Model(apiemployee.EmployeeInfo{}).Where("id = ? and is_delete = 1 ", employeeInfo.Id).Update(employeeInfo)
	return db.Error
}

func GetEmployeeInfo(id int64) *apiemployee.EmployeeInfo {
	employeeInfo := new(apiemployee.EmployeeInfo)
	coredb.Get().Model(apiemployee.EmployeeInfo{}).Where("id = ? and is_delete = 1 ", id).Find(employeeInfo)
	return employeeInfo
}

func AddEmployeeHealthInfo(employeeHealthInfo *apiemployee.EmployeeHealthInfo) error {
	db := coredb.Get().Model(apiemployee.EmployeeHealthInfo{}).Create(employeeHealthInfo)
	return db.Error
}

func UpdateEmployeeHealthInfo(employeeHealthInfo *apiemployee.EmployeeHealthInfo) error {
	db := coredb.Get().Model(apiemployee.EmployeeHealthInfo{}).Where("employee_id = ? and is_delete = 1 ", employeeHealthInfo.EmployeeId).Update(employeeHealthInfo)
	return db.Error
}

func GetEmployeeHealthInfo(employeeId int64) *apiemployee.EmployeeHealthInfo {
	employeeHealthInfo := new(apiemployee.EmployeeHealthInfo)
	coredb.Get().Model(apiemployee.EmployeeInfo{}).Where("employee_id = ?", employeeId).Find(employeeHealthInfo)
	return employeeHealthInfo
}

func GetEmployeeHealthInfos(employeeIds string) ([]*apiemployee.EmployeeHealthInfo, error) {
	empHealthInfos := make([]*apiemployee.EmployeeHealthInfo, 0)
	db := coredb.Get().Raw(fmt.Sprintf("select * from employee_health_infos where employee_id in (%s)", employeeIds)).Find(&empHealthInfos)

	if db.Error != nil {
		return nil, db.Error
	} else {
		return empHealthInfos, nil
	}
}

func AddEmployeeHealthRecord(employeeHealthRecord *apiemployee.EmployeeHealthRecord) error {
	db := coredb.Get().Model(apiemployee.EmployeeHealthRecord{}).Create(employeeHealthRecord)
	return db.Error
}

func CountGetEmployeeHealthRecordList(employeeId int64) int32 {
	var total int32
	coredb.Get().Model(apiemployee.EmployeeHealthRecord{}).Where(" employee_id = ? ", employeeId).Count(&total)
	return total
}

func GetEmployeeHealthRecordList(in *apiemployee.GetEmployeeHealthRecordListReq) []*apiemployee.EmployeeHealthRecord {
	employeeHealthRecords := make([]*apiemployee.EmployeeHealthRecord, 0)
	coredb.Get().Model(apiemployee.EmployeeHealthRecord{}).Where(" employee_id = ? ", in.EmployeeId).Find(&employeeHealthRecords)
	return employeeHealthRecords
}

func GetEmployeeHealthRecordListByCompanyId(companyId int64) []*apiemployee.EmployeeHealthRecord {
	sql := " SELECT ehr1.* FROM (SELECT * FROM employee_health_records ehr WHERE company_id = ? ) ehr1 LEFT JOIN (SELECT * FROM employee_health_records ehr WHERE company_id = ? ) ehr2 ON (ehr1.employee_id = ehr2.employee_id AND ehr1.id < ehr2.id ) WHERE ehr2.id IS NULL "
	employeeHealthRecords := make([]*apiemployee.EmployeeHealthRecord, 0)
	coredb.Get().Raw(sql, companyId, companyId).Scan(&employeeHealthRecords)
	return employeeHealthRecords
}

func GetEmployeeHealthRecordLast(empID int64) (*apiemployee.EmployeeHealthRecord, error) {
	e := &apiemployee.EmployeeHealthRecord{}

	db := coredb.Get().Raw("select * from employee_health_records where id = ? order by id desc limit 1", empID).Find(e)

	if db.Error != nil {
		return nil, db.Error
	}
	return e, nil
}

func GetStreetInfoTotal(streetID int64, query string) (int64, error) {
	var total int64
	buffer := new(bytes.Buffer)
	param := make([]interface{}, 0)
	buffer.WriteString(" select count(id) from employee_infos where is_delete = 1 and (company_street_id = ? or residence_street_id = ?)")
	param = append(param, streetID, streetID)
	if tool.IsNotBlank(query) {
		buffer.WriteString(" and (name like ? or phone like ?) ")
		param = append(param, "%"+query+"%", "%"+query+"%")
	}
	db := coredb.Get().Raw(buffer.String(), param...).Count(&total)
	if db.Error != nil {
		return 0, db.Error
	} else {
		return total, nil
	}

}

func GetStreetInfo(streetID int64, lastID int64, pageNum int32, pageSize int32, query string) ([]*apiemployee.EmployeeInfo, error) {
	e := make([]*apiemployee.EmployeeInfo, 0)
	var db *gorm.DB
	buffer := new(bytes.Buffer)
	param := make([]interface{}, 0)
	buffer.WriteString(" select * from employee_infos where is_delete = 1 and (company_street_id = ? or residence_street_id = ?) ")
	param = append(param, streetID, streetID)
	if tool.IsNotBlank(query) {
		buffer.WriteString(" and (name like ? or phone like ?) ")
		param = append(param, "%"+query+"%", "%"+query+"%")
	}
	buffer.WriteString(" order by id desc limit ?,? ")
	param = append(param, pageSize*(pageNum-1), pageSize)

	db = coredb.Get().Raw(buffer.String(), param...).Find(&e)

	if db.Error != nil {
		return nil, db.Error
	}
	return e, nil
}

func GetCommunityInfo(cID int64, lastID int64, pageNum int32, pageSize int32, query string) ([]*apiemployee.EmployeeInfo, error) {
	e := make([]*apiemployee.EmployeeInfo, 0)
	var db *gorm.DB
	buffer := new(bytes.Buffer)
	param := make([]interface{}, 0)
	buffer.WriteString("select * from employee_infos where is_delete = 1 and (company_community_id = ? or residence_community_id = ?)  ")
	param = append(param, cID, cID)
	if tool.IsNotBlank(query) {
		buffer.WriteString(" and (name like ? or phone like ?) ")
		param = append(param, "%"+query+"%", "%"+query+"%")
	}
	buffer.WriteString(" order by id desc limit ?,?")
	param = append(param, pageSize*(pageNum-1), pageSize)
	db = coredb.Get().Raw(buffer.String(), param...).Find(&e)

	if db.Error != nil {
		return nil, db.Error
	}
	return e, nil
}

func GetCommunityInfoTotal(cID int64, query string) (int64, error) {
	var total int64
	var db *gorm.DB
	buffer := new(bytes.Buffer)
	param := make([]interface{}, 0)
	buffer.WriteString("select count(id) from employee_infos where is_delete = 1 and (ompany_community_id = ? or residence_community_id = ?) ")
	param = append(param, cID, cID)
	if tool.IsNotBlank(query) {
		buffer.WriteString(" and (name like ? or phone like ?) ")
		param = append(param, "%"+query+"%", "%"+query+"%")
	}
	db = coredb.Get().Raw(buffer.String(), param...).Count(&total)
	if db.Error != nil {
		return 0, db.Error
	}
	return total, nil
}

func DeleteEmployInfo(empID string) error {
	db := coredb.Get().Exec(fmt.Sprintf("update employee_infos set is_delete = 2 where id in (%s)", empID))

	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("employeeID not exit")
	}

	return nil
}
