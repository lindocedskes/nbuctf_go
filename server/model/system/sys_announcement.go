package system

import "github.com/lindocedskes/model"

type Announcement struct {
	model.BaseModel
	Title   string `json:"title" gorm:"comment:标题"`             // 标题
	Content string `json:"content" gorm:"type:text;comment:内容"` // 内容
}
