package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	systemReq "github.com/lindocedskes/model/system/request"
	systemRes "github.com/lindocedskes/model/system/response"
	"github.com/lindocedskes/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"
)

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var u_l systemReq.Login       // 声明Login结构体
	err := c.ShouldBindJSON(&u_l) // 从请求中解析出Login结构体
	key := c.ClientIP()           // 获取客户端IP

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(u_l, utils.LoginVerify) // 验证器，utils.LoginVerify是规则集，
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否正确
	openCaptcha := global.NBUCTF_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆（最多）次数
	openCaptchaTimeOut := global.NBUCTF_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)                                   // 获取缓存
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut)) // 设置缓存，记录次数，超时时间
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v) // 判断是否开启验证码

	//store.Verify(u_l.CaptchaId, u_l.Captcha, true)  // 验证码验证
	if !oc || (u_l.CaptchaId != "" && u_l.Captcha != "" && store.Verify(u_l.CaptchaId, u_l.Captcha, true)) {
		u := &system.SysUser{Username: u_l.Username, Password: u_l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}
	// 验证码相关，记录请求次数+1，超上限要求验证码
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.NBUCTF_CONFIG.JWT.SigningKey)} // 唯一签名，
	claims := j.CreateClaims(systemReq.BaseClaims{                           //
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	//多点登录拦截
	if !global.NBUCTF_CONFIG.System.UseMultipoint { //是否启用
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix())) //设cookie中 "x-token" 字段值是 JWT 的字符串
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}
	//多点登录拦截 启用
	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil { //检查用户的 JWT 是否已经被存储在 Redis 中
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil { //已登录存在jwt
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		//不存在jwt
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
	}
}
