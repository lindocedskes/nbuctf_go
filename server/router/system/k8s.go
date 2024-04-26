package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type K8sRouter struct{}

// InitK8sRouter 初始化 k8s 路由
func (s *K8sRouter) InitK8sRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	k8sRouter := Router.Group("k8s")
	k8sRouterWithRecord := Router.Group("k8s").Use(middleware.OperationRecord()) //记录关键操作
	k8sApi := v1.ApiGroupApp.SystemApiGroup.K8sApi
	{
		k8sRouterWithRecord.POST("openContainer", k8sApi.OpenContainer)   // 开启容器
		k8sRouterWithRecord.POST("closeContainer", k8sApi.CloseContainer) // 关闭容器
		k8sRouterWithRecord.POST("checkContainer", k8sApi.CheckContainer) // 查询容器运行状态

		k8sRouterWithRecord.GET("closeAllContainer", k8sApi.CloseAllContainer) // 关闭所有容器
	}
	return k8sRouter

}
