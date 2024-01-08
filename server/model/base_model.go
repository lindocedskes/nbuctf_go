package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primaryKey" json:"id"` //struct tag，告诉 Gorm 这个字段是数据库表的主键，告诉 JSON 包在序列化和反序列化时，这个字段的名称应该是 id
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` //为该字段创建索引
}
