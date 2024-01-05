package model

import "gorm.io/gorm"

// 文章
type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `grom:"type:varchar(100);not null" json:"title"`
	Cid     int    `grom:"type:int;not null" json:"cid"`
	Desc    string `grom:"type:varchar(200)" json:"desc"`
	Content string `grom:"type:longtext" json:"content"`
	Img     string `grom:"type:varchar(100)" json:"img"`
}
