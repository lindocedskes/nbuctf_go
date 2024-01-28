package system

import (
	"github.com/lindocedskes/model"
)

type JwtBlacklist struct {
	model.BaseModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
