// 存放初始化数据的函数
package source

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lindocedskes/common"
	"github.com/lindocedskes/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sync"
)

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

type InitCasbin struct {
}

func GetCasbin() *casbin.CachedEnforcer {
	once.Do(func() { //确保函数只执行一次
		a, _ := gormadapter.NewAdapterByDB(common.NBUCTF_DB)
		//模型字符串
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act, eft
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow)) && !some(where (p.eft == deny))
		
		[matchers]
		m = (r.sub == p.sub || g(r.sub, p.sub)) && keyMatch(r.obj, p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text) // 从字符串加载模型
		if err != nil {
			fmt.Println("字符串加载模型失败!")
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a) //创建一个新的缓存Enforcer
		cachedEnforcer.SetExpireTime(60 * 60)              // 设置缓存过期时间
		_ = cachedEnforcer.LoadPolicy()                    //从数据库加载策略
	})
	return cachedEnforcer
}

// 创建gormadapter.CasbinRule 表
func (*InitCasbin) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&gormadapter.CasbinRule{})
}

// 返回 gormadapter.CasbinRule 表的名称
func (i InitCasbin) InitName() string {
	var entity gormadapter.CasbinRule
	return entity.TableName()
}

// 初始化 Casbin 的规则数据
func (i *InitCasbin) InitData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	rules := getDefaultCasbinRules() // 获取默认的 Casbin 规则数据
	if err := db.Create(&rules).Error; err != nil {
		return ctx, errors.Wrap(err, "Casbin 表 ("+i.InitName()+") 数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitName(), rules)
	return next, nil
}

// 检查特定的 Casbin 规则是否已经插入
func (i *InitCasbin) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where(gormadapter.CasbinRule{Ptype: "p", V0: "admin", V1: "/index", V2: "GET"}).
		First(&gormadapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

// 默认的 Casbin 规则数据
func getDefaultCasbinRules() []gormadapter.CasbinRule {
	rules := []gormadapter.CasbinRule{
		{Ptype: "p", V0: "admin", V1: "/index", V2: "GET", V3: "allow"},
		{Ptype: "p", V0: "admin", V1: "/auth/login", V2: "POST", V3: "allow"},
	}

	return rules
}
