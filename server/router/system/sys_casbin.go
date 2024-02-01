package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord()) //使用中间件：记录http操作日志
	casbinRouterWithoutRecord := Router.Group("casbin")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.CasbinApi //获取了CasbinApi对象，里面有Casbin相关的方法
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin) // 更新casbin权限
	}
	{
		casbinRouterWithoutRecord.POST("getCasbinRuleByAuthorityId", casbinApi.GetPolicyPathByAuthorityId) // 获取权限列表
	}
}
