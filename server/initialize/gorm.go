package initialize

import (
	"fmt"
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lindocedskes/model/system"
	sysModel "github.com/lindocedskes/model/system"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

// 创建（或更新）数据库表的结构
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate( //自动迁移（创建或更新）数据库表
		sysModel.SysUser{},       // 用户表
		gormadapter.CasbinRule{}, // 权限表
		sysModel.SysAuthority{},  // 角色权限表
		sysModel.JwtBlacklist{},  // jwt黑名单
		sysModel.SysHttpRecord{}, // http请求记录表
		sysModel.SysBaseMenu{},   // 菜单表

		system.SysBaseMenuParameter{}, // 菜单参数
		system.SysBaseMenuBtn{},       //	基础菜单按钮
		system.SysAuthorityBtn{},      // 角色按钮关联表

		sysModel.SysFileUploadAndDownload{}, // 文件上传-删除记录表
		sysModel.SysFile{},                  //文件切片上传结构存储
		sysModel.SysFileChunk{},

		sysModel.Question{},        // 题目表
		sysModel.RightGameRecord{}, // 正确题目记录表 题目id-用户id
		sysModel.WrongGameRecord{}, // 错误题目记录表 题目id-用户id
		//sysModel.QuestionFile{},    // 题目-文件关联表 -自动生成
		sysModel.UserScore{}, // 用户-分数表
	)

	if err != nil {
		os.Exit(0)
	}
	fmt.Println("数据表AutoMigrate完成")
}
