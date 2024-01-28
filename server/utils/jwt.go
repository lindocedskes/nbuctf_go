package utils

import (
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system/request"
	"time"
)

type JWT struct {
	SigningKey []byte
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(global.NBUCTF_CONFIG.JWT.SigningKey),
	}
}

// 创建JWT 声明类实例，包含相关配置
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.NBUCTF_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.NBUCTF_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.NBUCTF_CONFIG.JWT.Issuer,           // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //根据hash算法、配置声明进行hash签名
	return token.SignedString(j.SigningKey)                    //根据密钥，返回jwt字符串和err
}
