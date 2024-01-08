package source

import (
	"context"
	sysModel "github.com/lindocedskes/model/system"
	"github.com/lindocedskes/service/system"
	"github.com/lindocedskes/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type InitUser struct {
}

func (i *InitUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i *InitUser) InitName() string {
	return sysModel.SysUser{}.TableName()
}

func (i *InitUser) InitData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	user := getDefaultUser(ctx)
	if err := db.Create(&user).Error; err != nil {
		return ctx, errors.Wrap(err, "User 表 ("+i.InitName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitName(), user)
	return next, nil
}

func (i *InitUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	rootUser := db.Where("username = ?", "root").First(&sysModel.SysUser{})
	if errors.Is(rootUser.Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

func getDefaultUser(ctx context.Context) []sysModel.SysUser {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return []sysModel.SysUser{}
	}
	var rootRole sysModel.SysRole
	db.Where("role_name = ?", "root").First(&rootRole)
	root_pwd := utils.BcryptHash("root666!!!")
	root_secret := utils.RandomString(16)
	users := []sysModel.SysUser{
		{Username: "root", Password: root_pwd, Email: "", Roles: []sysModel.SysRole{rootRole}, Enable: 1, SecretKey: root_secret},
	}
	return users
}
