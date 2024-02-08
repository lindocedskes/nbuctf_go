package system

import (
	"context"
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lindocedskes/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/*管理 CasbinRule 表的初始化器*/

type initApi struct{}

const initOrderCasbin = initOrderApi + 1

type initCasbin struct{}

// auto run
func init() {
	system.RegisterInit(initOrderCasbin, &initCasbin{})
}

// 迁移表
func (i *initCasbin) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&adapter.CasbinRule{})
}

func (i *initCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&adapter.CasbinRule{})
}

func (i initCasbin) InitializerName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (i *initCasbin) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := getDefaultCasbinRule()
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitializerName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(adapter.CasbinRule{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"}).
		First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// 初始 CasbinRule 表数据 TODO add data
// Ptype: "p" 表示权限，Ptype: "g" 表示用户组，V0: "888" 表示角色ID，V1: "/user/getUserInfo" 表示接口路径，V2: "GET" 表示请求方法
func getDefaultCasbinRule() []adapter.CasbinRule {
	return []adapter.CasbinRule{
		//管理员的权限
		{Ptype: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/init/checkdb", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/register", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/changePassword", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/setUserAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/setUserAuthorities", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/deleteUser", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/setUserInfo", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/resetPassword", V2: "PUT"},
		{Ptype: "p", V0: "888", V1: "/userbyadmin/getUserList", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/createAuthority", V2: "POST"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/deleteAuthority", V2: "DELETE"},
		{Ptype: "p", V0: "888", V1: "/authoritybyadmin/getAuthorityList", V2: "POST"},
		//普通用户的权限
		{Ptype: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"},
	}
}
