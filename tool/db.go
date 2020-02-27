package tool

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// v5数据库参数对象
var robotdb DBParameters

func init() {
	// 初始化v5数据库参数
	robotdb.Address = "10.10.2.128"
	robotdb.UserName = "root"
	robotdb.Password = "root"
	robotdb.Port = "3306"
	robotdb.DB = "robot"
}

// 数据库参数
type DBParameters struct {
	Address  string //数据库地址
	UserName string //用户名
	Password string //密码
	DB       string //数据库
	Port     string // 端口
}

// 获取数据库连接
func GetDB() *gorm.DB {
	dataBaseConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", robotdb.UserName, robotdb.Password, robotdb.Address, robotdb.Port, robotdb.DB)
	db, _ := gorm.Open("mysql", dataBaseConfig)
	//开启sql调试模式
	db.LogMode(true)
	//连接池的空闲数大小
	db.DB().SetMaxIdleConns(5)
	//最大打开连接数
	db.DB().SetMaxOpenConns(5)
	db.DB().SetMaxOpenConns(5)
	return db
}
