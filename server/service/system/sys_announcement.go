package system

import (
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
)

type AnnouncementService struct{}

// GetAnnouncementList 获取公告列表
func (a *AnnouncementService) GetAnnouncementList() ([]system.Announcement, error) {
	db := global.NBUCTF_DB.Model(&system.Announcement{})
	var announcementList []system.Announcement
	err := db.Find(&announcementList).Error
	return announcementList, err
}

// CreateAnnouncement 创建公告
func (a *AnnouncementService) CreateAnnouncement(announcement system.Announcement) error {
	err := global.NBUCTF_DB.Create(&announcement).Error
	return err
}

// DeleteAnnouncement 删除公告 by id
func (a *AnnouncementService) DeleteAnnouncement(id uint) error {
	var announcement system.Announcement
	err := global.NBUCTF_DB.Where("id = ?", id).Delete(&announcement).Error
	return err
}

// UpdateAnnouncement 更新公告
func (a *AnnouncementService) UpdateAnnouncement(announcement system.Announcement) error {
	err := global.NBUCTF_DB.Model(&system.Announcement{}).Where("id = ?", announcement.ID).Updates(&announcement).Error //Updates 会忽略零值字段
	//err := global.NBUCTF_DB.Save(&announcement).Error //更新所有字段，包括空字段。

	return err
}
