package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authoritybyadmin").Use(middleware.OperationRecord())
	//authorityRouterWithoutRecord := Router.Group("authoritybyadmin")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.DELETE("deleteAuthority", authorityApi.DeleteAuthority) // 删除角色
	}
	{
		//authorityRouterWithoutRecord.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
	}
}
