package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
	"github.com/lindocedskes/middleware"
)

type AnnouncementRouter struct{}

func (s *AnnouncementRouter) InitAnnouncementRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	announcementRouter := Router.Group("announcement")
	announcementRouterWithRecord := Router.Group("announcementadmin").Use(middleware.OperationRecord()) //记录关键操作
	announcementApi := v1.ApiGroupApp.SystemApiGroup.AnnouncementApi
	{
		announcementRouter.GET("list", announcementApi.GetAnnouncementList)             // 获取公告列表
		announcementRouterWithRecord.POST("create", announcementApi.CreateAnnouncement) // 创建公告
		announcementRouterWithRecord.POST("delete", announcementApi.DeleteAnnouncement) // 删除公告
		announcementRouterWithRecord.POST("update", announcementApi.UpdateAnnouncement) // 更新公告
	}
	return announcementRouter
}
