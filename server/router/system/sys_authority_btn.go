package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
)

type AuthorityBtnRouter struct{}

// btn 细粒度控制，可以不用，SysBaseMenuBtn{} 未使用
func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router *gin.RouterGroup) {
	//authorityRouter := Router.Group("authorityBtn").Use(middleware.OperationRecord())
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouterWithoutRecord.POST("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)             // 获取权限按钮，按角色id & 菜单id 查询菜单按钮id
		authorityRouterWithoutRecord.POST("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)             // 重设权限按钮
		authorityRouterWithoutRecord.POST("canRemoveAuthorityBtn", authorityBtnApi.CanRemoveAuthorityBtn) // 查询是否能删除 菜单按钮id，在SysAuthorityBtn表中则不能
	}
}
