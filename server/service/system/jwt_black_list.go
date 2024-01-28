package system

import (
	"context"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
)

type JwtService struct{}

// 启动时加载
func LoadAll() {
	var data []string
	err := global.NBUCTF_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error // 从数据库加载jwt黑名单数据
	if err != nil {
		global.GVA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{}) // jwt黑名单 加入 BlackCache 中
	} // jwt黑名单 加入 BlackCache 中
}

// @description: 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.NBUCTF_REDIS.Get(context.Background(), userName).Result() //按userName
	return redisJWT, err
}

// @description: 保存jwt到redis
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	timer, err := utils.ParseDuration(global.NBUCTF_CONFIG.JWT.ExpiresTime) // 解析jwt过期时间
	if err != nil {
		return err
	}
	// 按userName，保存jwt到redis，上下文对象、键、值和过期时间
	err = global.NBUCTF_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

// @description: 拉黑jwt，添加到数据库和cache
func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	// Create 创建jwt黑名单 记录
	err = global.NBUCTF_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{}) //加入jwt作为键，值为空结构体
	return
}
