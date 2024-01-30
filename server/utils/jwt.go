package utils

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system/request"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
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

// CreateTokenByOldToken 创建新token 。用旧token标识 互斥锁避免并发问题。
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	//Do 方法首先获取 Group 的互斥锁,再执行，给定键一次只执行一次
	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	//令牌的字符串、空的 CustomClaims 结构体（用于存储解析出的声明）、函数（用于获取签名密钥）
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil { //解析出的token不为空
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid { //解析出的声明是CustomClaims类型，并且token有效
			return claims, nil //返回声明和err
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
