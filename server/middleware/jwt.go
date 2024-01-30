package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/service"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
	"strconv"
	"time"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

// 处理 JWT 认证，包括获取和解析令牌，处理令牌过期和黑名单，以及创建新的令牌。
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := utils.GetToken(c)
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			utils.ClearToken(c)
			c.Abort() //中止，不调用此请求的其余处理程序。
			return
		}
		j := utils.NewJWT() // 生成jwt实例
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil { //解析失败
			if errors.Is(err, utils.TokenExpired) {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				utils.ClearToken(c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			utils.ClearToken(c)
			c.Abort()
			return
		}
		//刷新和续期jwt
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime { //过期时间是否小于缓冲时间。如果是，那么它会创建一个新的令牌
			dr, _ := utils.ParseDuration(global.NBUCTF_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(dr))
			newToken, _ := j.CreateTokenByOldToken(token, *claims) //通过旧jwt创建新的令牌
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
			utils.SetToken(c, newToken, int(dr.Seconds())) //添加新的token到cookie
			if global.NBUCTF_CONFIG.System.UseMultipoint { //启用了多点登录
				RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username) //从redis中获取jwt
				if err != nil {
					global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken}) //把旧的jwt拉黑
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username) //把新的jwt存储到redis
			}
		}
		c.Set("claims", claims) //把解析出来的claims放到上下文中

		// 在每次请求被处理之前的逻辑

		c.Next() //决定是否继续执行后续的中间件或处理函数

		// 在请求被处理之后的逻辑
	}
}
