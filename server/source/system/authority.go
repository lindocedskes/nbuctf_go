package system

import (
	"context"
	sysModel "github.com/lindocedskes/model/system"
	"github.com/lindocedskes/service/system"
	"github.com/lindocedskes/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 管理 SysAuthority 表的初始化器
const initOrderAuthority = initOrderCasbin + 1 //当前初始化器的order顺序

type initAuthority struct{}
type initUser struct{}

// auto run，go特性，init()函数在每个包被导入时自动执行
func init() {
	system.RegisterInit(initOrderAuthority, &initAuthority{}) // 注册初始化器
}

func (i initAuthority) InitializerName() string {
	return sysModel.SysAuthority{}.TableName()
}

func (i *initAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuthority{})
}

func (i *initAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i *initAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysAuthority{
		{AuthorityId: 888, AuthorityName: "管理员", ParentId: utils.Pointer[uint](0), DefaultRouter: "game"},
		{AuthorityId: 777, AuthorityName: "队长", ParentId: utils.Pointer[uint](0), DefaultRouter: "game"},
		{AuthorityId: 7771, AuthorityName: "队员", ParentId: utils.Pointer[uint](777), DefaultRouter: "game"},
	}

	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!", sysModel.SysAuthority{}.TableName())
	}
	// data authority
	if err := db.Model(&entities[0]).Association("DataAuthorityId").Replace(
		[]*sysModel.SysAuthority{
			{AuthorityId: 888},
			{AuthorityId: 777},
			{AuthorityId: 7771},
		}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!",
			db.Model(&entities[0]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}
	if err := db.Model(&entities[1]).Association("DataAuthorityId").Replace(
		[]*sysModel.SysAuthority{
			{AuthorityId: 777},
			{AuthorityId: 7771},
		}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败!",
			db.Model(&entities[1]).Association("DataAuthorityId").Relationship.JoinTable.Name)
	}

	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("authority_id = ?", "888").
		First(&sysModel.SysAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
