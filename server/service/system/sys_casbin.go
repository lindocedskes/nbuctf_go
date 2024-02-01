package system

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system/request"
	"go.uber.org/zap"
	"strconv"
	"sync"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService) //创建了一个CasbinServiceApp对象实例

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer // 执行器（Enforcer），用于执行访问控制策略。使用了缓存和同步机制,线程安全
	once                 sync.Once                    //并发下只执行一次操作的机制
)

// @description: 根据数据库，规则模型 生成casbin执行器Enforcer
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

// @description: 更新casbin权限，角色id，路径，方法，先清除原有AuthorityID的再插入
func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	casbinService.ClearCasbin(0, authorityId) //清除sub为0，即角色id为authorityId的权限
	rules := [][]string{}
	//做权限去重处理
	deduplicateMap := make(map[string]bool) //创建一个map，key为string，value为bool
	for _, v := range casbinInfos {
		key := authorityId + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok { //如果key不存在，ok为false
			deduplicateMap[key] = true
			rules = append(rules, []string{authorityId, v.Path, v.Method}) //添加到rules中
		}
	}
	e := casbinService.Casbin()
	success, err := e.AddPolicies(rules) //添加casbin规则
	if err != nil {
		global.GVA_LOG.Error("添加casbin规则失败!", zap.Error(err))
	}
	if !success {
		return errors.New("存在相同casbin rule,添加失败h或其他错误查看日志")
	}
	return nil
}

// @description: 清除匹配的权限
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	//整数参数表示策略的字段索引，0为sub，字符串切片表示策略的字段值。
	success, _ := e.RemoveFilteredPolicy(v, p...) //删除匹配的权限
	return success
}

// 按AuthorityID查询casbin_rule 表：获取所有 sub 字段（也就是角色 ID）为 authorityId 的策略信息（路径和方法）
func (casbinService *CasbinService) GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin() //获取执行器
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId) //获取满足特定条件的策略
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1], //路径
			Method: v[2], //方法
		})
	}
	return pathMaps
}
