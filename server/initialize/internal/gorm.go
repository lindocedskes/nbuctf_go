package internal

import "gorm.io/gorm"

var Gorm = new(_gorm)

type _gorm struct { //私有
}

// 提供一个方法来获取已经配置好的 gorm.Config 实例，禁用外键约束
func (g *_gorm) Config() *gorm.Config {
	//字段的作用是在迁移数据库表时禁用外键约束,返回配置实例指针
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}
