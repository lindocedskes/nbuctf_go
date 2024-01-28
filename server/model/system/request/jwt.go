package request

import (
	"github.com/gofrs/uuid/v5"
	"github.com/golang-jwt/jwt/v4"
)

// BaseClaims ，包含用户的基本信息
type BaseClaims struct {
	UUID        uuid.UUID // 用户UUID
	ID          uint      // 用户ID
	Username    string    // 用户名
	NickName    string    // 昵称
	AuthorityId uint      // 角色ID
}

// CustomClaims ，包含用户的基本信息和jwt标准claims
type CustomClaims struct {
	BaseClaims
	BufferTime           int64 // 缓冲时间，表示在令牌过期后的一段时间内，用户仍然可以使用旧的令牌获取新的令牌
	jwt.RegisteredClaims       // jwt标准claims
}
