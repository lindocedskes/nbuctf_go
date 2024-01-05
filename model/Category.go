package model

import "gorm.io/gorm"

// 文章分类
type Category struct {
	gorm.Model
	//新版gorm添加
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}
