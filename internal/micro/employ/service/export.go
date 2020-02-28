package service

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"robot-server/api/apiexport"
	"robot-server/api/base"
	"robot-server/internal/core/conn"
	"robot-server/internal/core/errorcode"
	"robot-server/internal/micro/employ/empdb"
	"robot-server/tool"
	"strconv"
	"strings"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

const (
	/*****************************************************导出标签******************************************************/
	// 公司导出
	company = "company"
	// 员工导出
	employee = "employee"
	/*****************************************************导出标签******************************************************/

	/**************************************************七牛配置相关*****************************************************/
	// 七牛地址
	url    = "http://assets.xiaokayun.cn"
	access = "5bJMhEn4DSLNAJT-JiIw9rhmk8coOxMpVwGZoCRc"
	secret = "mSDqSTWRySYhEatdMuGlNGKFQLhYD4Ue97XYiSD3"
	/***************************************************七牛配置相关****************************************************/

	/************************************************文件上传进度状态****************************************************/
	// 状态 1.创建任务；2.数据写入文件；3文件上传；4.完成
	create = 1
	write  = 2
	upload = 3
	finish = 4
	/************************************************文件上传进度状态****************************************************/

	/************************************************需要格式化的字段****************************************************/
	// 创建时间
	creatTime = "create_time"
	// 修改时间
	updateTime = "update_time"
	/************************************************需要格式化的字段****************************************************/
	// 每个csv文件承载数据条数
	limit = 1000
)

type ExportServer struct{}

// 导出通用方法
func (e *ExportServer) Export(ctx context.Context, in *apiexport.ExportReq) (*apiexport.ExportRes, error) {
	if in == nil {
		return &apiexport.ExportRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Tag) {
		return &apiexport.ExportRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	if tool.IsBlank(in.Field) {
		return &apiexport.ExportRes{Base: &base.BaseRes{ErrorCode: errorcode.ParameterError}}, nil
	}
	displays, filed := field(in.Field)
	// 拼装统计sql
	var countSql string
	var querySql string
	// 拼装查询sql
	tag := in.Tag
	switch tag {
	case company:
		countSql = "SELECT COUNT(id) FROM company_infos WHERE is_delete = 1"
		buffer := new(bytes.Buffer)
		if in.ProvinceId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND province_id = %d", in.ProvinceId))
		}
		if in.CityId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND city_id = %d", in.CityId))
		}
		if in.DistrictId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND district_id = %d", in.DistrictId))
		}
		if in.StreetId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND street_id = %d", in.StreetId))
		}
		if in.CommunityId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND community_id = %d", in.CommunityId))
		}
		querySql = fmt.Sprintf("SELECT %s FROM company_infos  LEFT JOIN (SELECT p1.* FROM prevention_records p1 LEFT JOIN prevention_records p2ON (p1.company_id = p2.company_id AND p1.id < p2.id)WHERE p2.id IS NULL) p ON company_infos.id = p.company_id WHERE is_delete = 1 %s  LIMIT %s,%s", filed, buffer.String(), "%d", "%d")
	case employee:
		countSql = "SELECT COUNT(id) FROM employee_infos"
		buffer := new(bytes.Buffer)
		if in.ProvinceId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND (company_province_id = %d OR residence_province_id = %d)", in.ProvinceId, in.ProvinceId))
		}
		if in.CityId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND (company_city_id = %d OR residence_city_id = %d)", in.CityId, in.CityId))
		}
		if in.DistrictId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND (company_district_id = %d OR residence_district_id = %d)", in.DistrictId, in.DistrictId))
		}
		if in.StreetId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND (company_street_id = %d OR residence_street_id = %d)", in.StreetId, in.StreetId))
		}
		if in.CommunityId != 0 {
			buffer.WriteString(fmt.Sprintf(" AND (company_community_id = %d OR residence_community_id = %d)", in.CommunityId, in.CommunityId))
		}
		querySql = fmt.Sprintf("SELECT %s FROM (SELECT ei.* , ehi.symptom_within14_days,ehi.form_hu_bei,ehi.touch_ncp,ehi.been_to_hu_bei,ehi.return_cheng_du_time,ehi.return_cheng_du_detailed_address,ehi.return_cheng_du_transport,ehi.transport_detail_info,ehi.touch_hu_bei_people_within14_days,ehi.last_touch_time,ehi.from_cheng_du_outside,ehi.this_place_time,ehi.return_this_place_transport,ehi.profession,ehi.life_trace,ehi.sign_image,ehi.come_from_wenjiang,ehi.come_from_address,ehi.come_from_detailed_address FROM employee_infos ei LEFT JOIN employee_health_infos ehi ON ei.id = ehi.employee_id WHERE is_delete = 1 %s LIMIT %s,%s) e LEFT JOIN (SELECT ehr1.employee_id, ehr1.temperature FROM employee_health_records ehr1 LEFT JOIN employee_health_records ehr2 ON (ehr1.employee_id = ehr2.employee_id AND ehr1.id < ehr2.id) WHERE ehr2.id IS NULL) eh ON e.id = eh.employee_id", filed, buffer.String(), "%d", "%d")
	default:
		return &apiexport.ExportRes{Base: &base.BaseRes{ErrorCode: errorcode.ExportTagError}}, nil
	}
	exportRecord := &apiexport.ExportRecord{
		Tag:      tag,
		CountSql: countSql,
		QuerySql: querySql,
		Status:   create,
	}
	// 添加导出开始记录
	err := empdb.AddExportRecord(exportRecord)
	if err != nil {
		return &apiexport.ExportRes{Base: &base.BaseRes{ErrorCode: errorcode.InsertError}}, nil
	}
	// 异步进行导出操作
	go func(exportRecord *apiexport.ExportRecord, displays []string, in *apiexport.ExportReq) {
		do(exportRecord, displays)
		// t推送前台
		byConn := conn.GetConnByConn(in.Head.ConnID)
		if byConn == nil {
			return
		}
		bytes, _ := json.Marshal(exportRecord)
		byConn.WriteMessage_v2(&base.Msg{ApiId: base.Api_ExportFinish, Data: bytes})
	}(exportRecord, displays, in)
	return &apiexport.ExportRes{Base: &base.BaseRes{ErrorCode: errorcode.Succeed}}, nil
}

// 前端传来的字段json字符串处理
func field(field string) ([]string, string) {
	fields := make([]*apiexport.Field, 0)
	json.Unmarshal([]byte(field), &fields)
	// 查询字段拼接
	var filed string
	displays := make([]string, 0)
	for index, fieldItem := range fields {
		// 时间戳格式化
		if fieldItem.Filed == creatTime {
			fieldItem.Filed = "(case when create_time != 0 then from_unixtime(create_time) else '' end) AS create_time"
		}
		if fieldItem.Filed == updateTime {
			fieldItem.Filed = "(case when update_time != 0 then from_unixtime(update_time) else '' end) AS update_time"
		}
		if index == 0 {
			filed = fieldItem.Filed
		}
		filed = fmt.Sprintf("%s%s%s", filed, ",", fieldItem.Filed)
		displays = append(displays, fieldItem.Display)
	}
	return displays, filed
}

// 异步导出动作
func do(exportRecord *apiexport.ExportRecord, displays []string) {
	// 修改任务为数据写入
	exportRecord.Status = write
	err := empdb.UpdateExportRecord(exportRecord)
	if err != nil {
		return
	}
	// 统计导出数据量
	total := empdb.CountExportData(exportRecord.CountSql)
	//取余（总数除以每个文件写入条数）
	fileNum := total % limit
	//获取生成文件数
	if fileNum == 0 {
		fileNum = total / limit
	} else {
		fileNum = total/limit + 1
	}
	//检查导出文件夹是否存在，如不存在创建文件夹
	filePath := "D:\\export\\"
	check, err := pathExist(filePath)
	if !check && err == nil {
		os.Mkdir(filePath, 0)
	}
	var fileList string
	//循环取数据写文件
	for i := 0; i < int(fileNum); i++ {
		//根据时间戳命名文件
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + ".csv"
		fileExist, err := pathExist(filePath + fileName)
		if fileExist && err == nil {
			//文件已存在，生成一个重命名文件
			fileName = strconv.FormatInt(time.Now().Unix(), 10) + "(" + strconv.Itoa(i) + ")" + ".csv"
		}
		f, err := os.Create(filePath + fileName) //创建文件
		if err != nil {
			return
		}
		fileList += fileName + ","
		//定义写出数组
		var writeData [][]string
		//写入列头
		writeData = append(writeData, displays)
		// 本次查询偏移量
		offset := i * limit
		//获取数据
		exportRecord.QuerySql = fmt.Sprintf(exportRecord.QuerySql, limit, offset)
		rows, err := empdb.GetExportData(exportRecord)
		if err != nil {
			return
		}
		cols, err := rows.Columns() // Remember to check err afterwards
		scanArgs := make([]interface{}, len(cols))
		values := make([]sql.RawBytes, len(cols))
		for i, _ := range cols {
			scanArgs[i] = &values[i]
		}
		for rows.Next() {
			err = rows.Scan(scanArgs...)
			var data []string
			for _, v := range values {
				data = append(data, fmt.Sprintf("%s%s", v, "\t"))
			}
			writeData = append(writeData, data)
		}
		rows.Close()

		bomUtf8 := []byte{0xEF, 0xBB, 0xBF}
		f.Write(bomUtf8)
		//创建一个新的写入文件流
		w := csv.NewWriter(f)
		//写出文件
		w.WriteAll(writeData)
		w.Flush()
		//关闭写入流
		f.Close()
	}
	fmt.Println("======================开始压缩文件夹===========================", fileList)
	//压缩文件
	fileList = subString(fileList, 0, len(fileList)-1)
	fileStr := strings.Split(fileList, ",")
	var files []*os.File
	for _, v := range fileStr {
		file, err := os.Open(filePath + v)
		if err != nil {
			panic(err)
		}
		files = append(files, file)
	}
	fileName := fmt.Sprintf("%d_%d", exportRecord.Id, time.Now().Unix()) + ".zip"
	dest := filePath + fileName
	zipper := doCompress(files, dest)
	if zipper != nil {
		fmt.Println("zip error:", zipper)
		return
	}
	fmt.Println("======================压缩完成=================================")
	fmt.Println("======================上传文件=================================")
	// 分文件夹上传
	if err := uptoqiniu(dest, fmt.Sprintf("export/%s", fileName)); err != nil {
		return
	}
	fmt.Println("======================更新任务状态和文件下载路径================")
	exportRecord.FilePath = fmt.Sprintf("%s%s%s", url, "/export/", fileName)
	exportRecord.Status = upload
	err = empdb.UpdateExportRecord(exportRecord)
	if err != nil {
		return
	}
	fmt.Println("======================删除文件=================================")
	for _, k := range fileStr {
		os.RemoveAll(filePath + k)
	}
	os.RemoveAll(dest)
	fmt.Println("======================推送消息给前台=================================")

	fmt.Println("======================导出完成=================================")
	exportRecord.Status = finish
	empdb.UpdateExportRecord(exportRecord)
}

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func doCompress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func uptoqiniu(localFile, key string) error {
	bucket := "images"
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 3600 * 24 * 365 * 2, //两年有效期
	}
	mac := qbox.NewMac(access, secret)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	resumeUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{}
	err := resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println("PutFile err :", err)
		return err
	} else {
		fmt.Println("PutFile success :", upToken)
	}
	fmt.Println(ret.Key, ret.Hash)
	upToken = ""
	return nil
}
