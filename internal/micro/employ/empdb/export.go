package empdb

import (
	"database/sql"
	"myserver/api/apiexport"
	"myserver/internal/core/coredb"
)

func AddExportRecord(exportRecord *apiexport.ExportRecord) error {
	db := coredb.Get().Model(apiexport.ExportRecord{}).Create(exportRecord)
	return db.Error
}

func UpdateExportRecord(exportRecord *apiexport.ExportRecord) error {
	db := coredb.Get().Model(apiexport.ExportRecord{}).Where("id = ?", exportRecord.Id).Update(exportRecord)
	return db.Error
}

func CountExportData(countSql string) int32 {
	var total int32
	coredb.Get().Exec(countSql).Count(&total)
	return total
}

func GetExportData(exportRecord *apiexport.ExportRecord) (*sql.Rows, error) {
	rows, err := coredb.Get().Raw(exportRecord.QuerySql).Rows()
	return rows, err
}
