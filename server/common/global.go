package common

import (
	"github.com/casbin/casbin/v2"
	"github.com/lindocedskes/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	NBUCTF_CONFIG config.Server
	NBUCTF_VP     *viper.Viper
	NBUCTF_DB     *gorm.DB
	NBUCTF_REDIS  *redis.Client

	NBUCTF_CASBIN *casbin.CachedEnforcer
)
