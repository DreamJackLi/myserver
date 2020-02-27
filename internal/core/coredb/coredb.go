package coredb

import (
	"fmt"
	"robot-server/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GORM配置信息
type GormConfig struct {
	// 方言
	Dialect string
	// 连接参数
	DialectArgs interface{}
	// 是否开启DEBUG(耗性能)
	IsDebug bool
	// 连接池的空闲数大小
	MaxIdleConns int
	// 最大打开连接数
	MaxOpenConns int
	// 连接最大存活时间
	ConnMaxLifetime time.Duration
}

var gormDB *gorm.DB

func Init() error {

	// 初始化数据库配置
	cfg := config.Get().Mysql
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		cfg.UserName,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DataBase,
	)
	gormConf := &GormConfig{
		Dialect:      "mysql",
		DialectArgs:  args,
		MaxIdleConns: int(cfg.MaxIdleConn),
		MaxOpenConns: int(cfg.MaxOpenConn),
		IsDebug:      cfg.Debug,
	}

	_, err := gormOpenDB(gormConf, nil)

	if err != nil {
		return err
	}

	return nil
}

func Get() *gorm.DB {
	return gormDB
}

// 开启Gorm数据库
func gormOpenDB(config *GormConfig, autoMigrates ...interface{}) (db *gorm.DB, err error) {
	db, err = gorm.Open(config.Dialect, config.DialectArgs)
	if err != nil {
		return
	}
	// 开启sql调试模式
	db.LogMode(config.IsDebug)
	// 连接池的空闲数大小
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	// 最大打开连接数
	db.DB().SetMaxOpenConns(config.MaxOpenConns)
	// 连接最大存活时间
	db.DB().SetConnMaxLifetime(config.ConnMaxLifetime)
	// // 自动建表
	// if autoMigrates != nil && len(autoMigrates) > 0 {
	//   db.AutoMigrate(autoMigrates...)
	// }
	gormDB = db
	return
}
