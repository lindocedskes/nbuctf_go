package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	//menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.CreateBaseMenu)        // 新增菜单
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系 -连接表增加记录
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // 更新菜单
	}
	{

	}
	return menuRouter
}
