package system

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lindocedskes/global"
	"go.uber.org/zap"
	"sync"
)

type CasbinService struct{}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer // 执行器（Enforcer），用于执行访问控制策略。使用了缓存和同步机制,线程安全
	once                 sync.Once                    //并发下只执行一次操作的机制
)

// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.Enforcer
func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() { //初始化代码只执行一次
		a, err := gormadapter.NewAdapterByDB(global.NBUCTF_DB) // 适配器，连接数据库和 Casbin
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		// Casbin 的RBAC模型 定义了请求、策略、角色、策略效果和匹配器的格式。
		//keyMatch2 内置函数 一个URL 路径或 : 模式下
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text) //创建了一个新的模型 实例m
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a) //*创建了一个新的执行器实例，根据模型和适配器
		syncedCachedEnforcer.SetExpireTime(60 * 60)                    //设置缓存过期时间1h
		_ = syncedCachedEnforcer.LoadPolicy()                          //从适配器加载策略
	})
	return syncedCachedEnforcer //返回执行器
}
