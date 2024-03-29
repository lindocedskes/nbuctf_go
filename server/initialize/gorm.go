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
		//sysModel.SysRole{},
		sysModel.SysApi{},        //存储api信息表
		sysModel.SysUser{},       // 用户表
		gormadapter.CasbinRule{}, // 权限表
		sysModel.SysAuthority{},  // 角色权限表
		sysModel.JwtBlacklist{},  // jwt黑名单
		sysModel.SysHttpRecord{}, // http请求记录表

		sysModel.SysBaseMenu{},          // 菜单表
		sysModel.SysBaseMenuParameter{}, // 存储菜单携带的参数（不必须）
		sysModel.SysBaseMenuBtn{},       // 角色菜单按钮关联表
		sysModel.SysAuthorityBtn{},      // 角色-按钮关联表,表示每个角色对每个菜单的哪些按钮有操作权限，实现更细粒度的权限控制

		sysModel.SysFileUploadAndDownload{}, // 文件上传-删除记录表
		sysModel.SysFile{},                  //文件切片上传结构存储
		sysModel.SysFileChunk{},
	)

	if err != nil {
		os.Exit(0)
	}
	fmt.Println("数据表AutoMigrate完成")
}
