package initialize

import (
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/utils"
	"github.com/songzhibin97/gkit/cache/local_cache"
)

// cache 配置初始化
func OtherInit() {
	dr, err := utils.ParseDuration(global.NBUCTF_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.NBUCTF_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache( //创建了一个新的本地缓存实例，
		local_cache.SetDefaultExpire(dr), //设置默认过期时间，根据配置文件中的JWT.ExpiresTime
	)
}
