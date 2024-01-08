package source

//初始化角色role表 数据
import (
	"context"
	sysModel "github.com/lindocedskes/model/system"
	"github.com/lindocedskes/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type InitRole struct {
}

// 判断是否存在表
func (i *InitRole) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysRole{})
}

func (i *InitRole) InitName() string {
	return sysModel.SysRole{}.TableName()
}

// 初始化数据
func (i *InitRole) InitData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	roles := getDefaultRoles() // 获取默认的角色数据
	if err := db.Create(&roles).Error; err != nil {
		return ctx, errors.Wrap(err, "Role 表 ("+i.InitName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitName(), roles)
	return next, nil
}

// 判断是否存在数据
func (i *InitRole) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	rootRole := db.Where("role_name = ?", "root").First(&sysModel.SysRole{})
	if errors.Is(rootRole.Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// 默认的角色数据
func getDefaultRoles() []sysModel.SysRole {
	roles := []sysModel.SysRole{
		{RoleName: "root"},
		{RoleName: "admin"},
		{RoleName: "normal"},
	}

	return roles
}
