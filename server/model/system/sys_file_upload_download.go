package system

import "github.com/lindocedskes/model"

type SysFileUploadAndDownload struct {
	model.BaseModel
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

func (SysFileUploadAndDownload) TableName() string {
	return "sys_file_upload_and_downloads"
}
