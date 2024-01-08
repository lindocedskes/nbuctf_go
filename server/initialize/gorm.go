package initialize

import (
	"fmt"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	sysModel "github.com/lindocedskes/model/system"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// 创建（或更新）数据库表的结构
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate( //自动迁移（创建或更新）数据库表
		sysModel.SysRole{},
		sysModel.SysUser{},
		gormadapter.CasbinRule{},
	)

	if err != nil {
		os.Exit(0)
	}
	fmt.Println("数据表AutoMigrate完成")
}
