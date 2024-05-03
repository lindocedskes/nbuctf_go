package initialize

import (
	"database/sql"
	"fmt"

	"github.com/lindocedskes/model/system/request"
	"github.com/lindocedskes/service"

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
	CreateDatabaseIfNotExist() //启用自动创建数据库
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
		global.NBUCTF_CONFIG.Mysql.Path,
		global.NBUCTF_CONFIG.Mysql.Port))

	if err != nil {
		global.GVA_LOG.Error("open database error", zap.Any("err", err))
	}
	//创建数据库如果不存在
	// _, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + global.NBUCTF_CONFIG.Mysql.Dbname)

	// fmt.Println("Host:", global.NBUCTF_CONFIG.Mysql.Path)
	// fmt.Println("Port:", global.NBUCTF_CONFIG.Mysql.Port)
	// if err != nil {
	// 	global.GVA_LOG.Error("create database error", zap.Any("err", err))
	// }

	// global.GVA_LOG.Info("成功新建数据库：" + global.NBUCTF_CONFIG.Mysql.Dbname)

	// Check if the database exists
	rows, err := db.Query("SHOW DATABASES LIKE '" + global.NBUCTF_CONFIG.Mysql.Dbname + "'")
	if err != nil {
		global.GVA_LOG.Error("check database error", zap.Any("err", err))
	}

	var dbName string
	if rows.Next() {
		rows.Scan(&dbName)
	}
	if dbName != global.NBUCTF_CONFIG.Mysql.Dbname {
		//数据库不存在
		if err == nil {
			// Database does not exist, 调用 InitDB 建库，建表，建数据
			err = service.ServiceGroupApp.SystemServiceGroup.InitDBService.InitDB(request.InitDB{
				DBType:   "mysql",
				Host:     global.NBUCTF_CONFIG.Mysql.Path,
				Port:     global.NBUCTF_CONFIG.Mysql.Port,
				UserName: global.NBUCTF_CONFIG.Mysql.Username,
				Password: global.NBUCTF_CONFIG.Mysql.Password,
				DBName:   global.NBUCTF_CONFIG.Mysql.Dbname,
			})

			if err != nil {
				global.GVA_LOG.Error("数据库初始化失败", zap.Error(err))
			}
		}
		global.GVA_LOG.Info("成功新建数据库：" + global.NBUCTF_CONFIG.Mysql.Dbname)
	} else {
		global.GVA_LOG.Info("数据库已存在：" + global.NBUCTF_CONFIG.Mysql.Dbname)
	}

}
