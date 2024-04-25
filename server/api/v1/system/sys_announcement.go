package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	"go.uber.org/zap"
)

type AnnouncementApi struct{}

// AnnouncementList 获取公告列表
func (a *AnnouncementApi) GetAnnouncementList(c *gin.Context) {
	AnnouncementList, err := announcementService.GetAnnouncementList()
	if err != nil {
		global.GVA_LOG.Error("获取公告列表失败!", zap.Error(err))
		response.FailWithMessage("获取公告列表失败", c)
		return
	}
	response.OkWithDetailed(AnnouncementList, "获取公告列表成功", c)
}

// CreateAnnouncement 创建公告
func (a *AnnouncementApi) CreateAnnouncement(c *gin.Context) {
	var announcement system.Announcement
	_ = c.ShouldBindJSON(&announcement)
	announcement.ID = 0 //创建公告时不需要ID
	if err := announcementService.CreateAnnouncement(announcement); err != nil {
		global.GVA_LOG.Error("创建公告失败!", zap.Error(err))
		response.FailWithMessage("创建公告失败", c)
		return
	}
	response.OkWithMessage("创建公告成功", c)
}

// DeleteAnnouncement 删除公告 by id
func (a *AnnouncementApi) DeleteAnnouncement(c *gin.Context) {
	var announcement system.Announcement
	_ = c.ShouldBindJSON(&announcement)
	if announcement.ID == 0 {
		response.FailWithMessage("删除失败,id请求错误", c)
		return
	}
	if err := announcementService.DeleteAnnouncement(announcement.ID); err != nil {
		global.GVA_LOG.Error("删除公告失败!", zap.Error(err))
		response.FailWithMessage("删除公告失败", c)
		return
	}
	response.OkWithMessage("删除公告成功", c)
}

// UpdateAnnouncement 更新公告
func (a *AnnouncementApi) UpdateAnnouncement(c *gin.Context) {
	var announcement system.Announcement
	_ = c.ShouldBindJSON(&announcement)
	if err := announcementService.UpdateAnnouncement(announcement); err != nil {
		global.GVA_LOG.Error("更新公告失败!", zap.Error(err))
		response.FailWithMessage("更新公告失败", c)
		return
	}
	response.OkWithMessage("更新公告成功", c)
}
