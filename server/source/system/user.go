package system

import (
	"context"
	"github.com/gofrs/uuid/v5"
	sysModel "github.com/lindocedskes/model/system"
	"github.com/lindocedskes/service/system"
	"github.com/lindocedskes/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderUser = initOrderAuthority + 1

type InitUser struct {
}

// auto run，go特性，init()函数在每个包被导入时自动执行
func init() {
	system.RegisterInit(initOrderUser, &initUser{})
}

// 迁移表
func (i *initUser) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysUser{})
}

func (i *initUser) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysUser{})
}

func (i *initUser) InitializerName() string {
	return sysModel.SysUser{}.TableName()
}

// 创建用户并设置用户的权限
func (i *initUser) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	user := getDefaultUser() // 获取默认用户数据
	if err := db.Create(&user).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysUser{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), user) // 将用户信息存入上下文
	//关联用户的权限表
	authorityEntities, ok := ctx.Value(initAuthority{}.InitializerName()).([]sysModel.SysAuthority)
	if !ok {
		return next, errors.Wrap(system.ErrMissingDependentContext, "创建 [用户-权限] 关联失败, 未找到权限表初始化数据")
	}
	if err := db.Model(&user[0]).Association("Authorities").Replace(authorityEntities); err != nil {
		return next, err
	}
	if err := db.Model(&user[1]).Association("Authorities").Replace(authorityEntities[:1]); err != nil {
		return next, err
	}
	return next, nil
}

// 查询是否存在admin用户
func (i *initUser) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	var record sysModel.SysUser
	if errors.Is(db.Where("username = ?", "a303176530").
		Preload("Authorities").First(&record).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return len(record.Authorities) > 0 && record.Authorities[0].AuthorityId == 888
}

// 初始化 root用户 数据
func getDefaultUser() []sysModel.SysUser {
	root_pwd := utils.BcryptHash("123456")

	users := []sysModel.SysUser{ // 初始化 root用户 数据
		{
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    "admin",
			Password:    root_pwd,
			NickName:    "管理员",
			HeaderImg:   "",
			AuthorityId: 888,
			Phone:       "",
			Email:       "",
		},
		{
			UUID:        uuid.Must(uuid.NewV4()),
			Username:    "user",
			Password:    root_pwd,
			NickName:    "用户1",
			HeaderImg:   "",
			AuthorityId: 9528,
			Phone:       "",
			Email:       "",
		},
	}
	return users
}
