package initialize

import (
	"database/sql"
	"fmt"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/initialize/internal"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.NBUCTF_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	CreateDatabaseIfNotExist()
	if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config()); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// 若数据库不存在，则创建数据库
func CreateDatabaseIfNotExist() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		global.NBUCTF_CONFIG.Mysql.Username,
		global.NBUCTF_CONFIG.Mysql.Password,
		global.NBUCTF_CONFIG.Mysql.Host,
		global.NBUCTF_CONFIG.Mysql.Port))

	if err != nil {
		global.GVA_LOG.Error("open database error", zap.Any("err", err))
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + global.NBUCTF_CONFIG.Mysql.Dbname)

	fmt.Println("Host:", global.NBUCTF_CONFIG.Mysql.Host)
	fmt.Println("Port:", global.NBUCTF_CONFIG.Mysql.Port)
	if err != nil {
		global.GVA_LOG.Error("create database error", zap.Any("err", err))
	}

	global.GVA_LOG.Info("成功新建数据库：" + global.NBUCTF_CONFIG.Mysql.Dbname)
}
