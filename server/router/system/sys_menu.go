package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord())
	menuRouterWithoutRecord := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("addBaseMenu", authorityMenuApi.CreateBaseMenu)        // 新增菜单
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系 -连接表增加记录
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateBaseMenu)     // 更新菜单
	}
	{
		menuRouterWithoutRecord.GET("getMenu", authorityMenuApi.GetMenu)                    // 获取的是与特定用户权限相关的菜单树
		menuRouterWithoutRecord.GET("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)    // 获取的是所有的基础菜单树
		menuRouterWithoutRecord.POST("getMenuList", authorityMenuApi.GetMenuList)           // 分页获取基础menu列表
		menuRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色id对应菜单
		menuRouterWithoutRecord.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单
	}
	return menuRouter
}
