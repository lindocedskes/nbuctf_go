package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/lindocedskes/config"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	NBUCTF_CONFIG config.Server
	NBUCTF_VP     *viper.Viper
	NBUCTF_DB     *gorm.DB
	NBUCTF_REDIS  *redis.Client

	NBUCTF_CASBIN *casbin.CachedEnforcer

	// GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	//GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
)
