package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model"
	"github.com/lindocedskes/model/common/request"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	systemReq "github.com/lindocedskes/model/system/request"
	systemRes "github.com/lindocedskes/model/system/response"
	"github.com/lindocedskes/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// @Summary  用户登录
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

// @Summary  管理员注册用户账号，需要管理员权限，普通用户注册不是这个
func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority //角色组
	for _, v := range r.AuthorityIds {    //遍历一个用户的所有角色id
		//一个用户的每有一个角色，都转化为 system.SysAuthority 实例记录AuthorityId，然后添加到 authorities 切片中
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, AuthorityId: r.AuthorityId, Authorities: authorities, Enable: r.Enable, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user) //注册，增
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// @Summary   用户注册账号
func (b *BaseApi) RegisterUser(c *gin.Context) {
	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	key := c.ClientIP() // 获取客户端IP

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.RegisterUserVerify)
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
	if !oc || (r.CaptchaId != "" && r.Captcha != "" && store.Verify(r.CaptchaId, r.Captcha, true)) {
		var authorities []system.SysAuthority //角色组
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: 777,
		})
		//固定角色id
		user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, AuthorityId: 777, Authorities: authorities, Phone: r.Phone, Email: r.Email}
		userReturn, err := userService.Register(*user)
		if err != nil {
			global.GVA_LOG.Error("注册失败!", zap.Error(err))
			response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败,"+err.Error()+"请换一个用户名", c)
			return
		}
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
		return
	}

	// 验证码相关，记录请求次数+1，超上限要求验证码
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// /@Summary   修改用户密码
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c) //从Gin的Context中获取从jwt解析出来的用户ID，避免越权
	u := &system.SysUser{BaseModel: model.BaseModel{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// @Summary   更改用户权限，修改主角色ID + 重设token
func (b *BaseApi) SetUserAuthority(c *gin.Context) {
	var sua systemReq.SetUserAuth
	err := c.ShouldBindJSON(&sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), c)
		return
	}
	userID := utils.GetUserID(c)                                //从Gin的Context中获取从jwt解析出来的用户ID，避免越权
	err = userService.SetUserAuthority(userID, sua.AuthorityId) //更新用户的主角色ID
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	//重新生成token，并将其设置为响应的头部和设置到cookie中
	claims := utils.GetUserInfo(c)                                           //从请求的上下文获取 已经jwt字符串解析还原的Claims
	j := &utils.JWT{SigningKey: []byte(global.NBUCTF_CONFIG.JWT.SigningKey)} // 唯一签名
	claims.AuthorityId = sua.AuthorityId                                     //修改后的角色ID
	token, err := j.CreateToken(*claims)                                     //重新生成token
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		c.Header("new-token", token) //设置响应头部
		c.Header("new-expires-at", strconv.FormatInt(claims.ExpiresAt.Unix(), 10))
		utils.SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60)) //设置到cookie中
		response.OkWithDetailed(struct {                                              //匿名结构体
			AuthorityId uint   `json:"authorityId"`
			Token       string `json:"token"`
			ExpiresAt   int64  `json:"expiresAt"`
		}{
			AuthorityId: sua.AuthorityId,
			Token:       token,
			ExpiresAt:   (claims.ExpiresAt.Unix() - time.Now().Unix()) / 60,
		}, "修改成功", c)
	}
}

// @Summary   设置用户权限-修改用户组角色id，并将角色组第一个设为主角色（AuthorityId）
func (b *BaseApi) SetUserAuthorities(c *gin.Context) {
	var sua systemReq.SetUserAuthorities
	err := c.ShouldBindJSON(&sua)
	fmt.Println("sua:", sua)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.SetUserAuthorities(sua.ID, sua.AuthorityIds)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// @Summary   删除用户，按id 删除对应用户，不能删除自身
func (b *BaseApi) DeleteUser(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	jwtId := utils.GetUserID(c)
	if jwtId == uint(reqId.ID) {
		response.FailWithMessage("删除失败, 不能删除自身", c)
		return
	}
	err = userService.DeleteUser(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// @Summary   设置用户信息，指定用户id
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify) //验证id
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(user.AuthorityIds) != 0 { //如果角色组id不为空
		err = userService.SetUserAuthorities(user.ID, user.AuthorityIds) //修改用户角色组，并更新第一个角色为主角
		if err != nil {
			global.GVA_LOG.Error("设置失败!", zap.Error(err))
			response.FailWithMessage("设置失败", c)
			return
		}
	}
	err = userService.SetUserInfo(system.SysUser{
		BaseModel: model.BaseModel{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// @Summary   设置自己的用户信息，从jwt解析出来的自身的用户ID
func (b *BaseApi) SetSelfInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user.ID = utils.GetUserID(c)
	err = userService.SetSelfInfo(system.SysUser{
		BaseModel: model.BaseModel{
			ID: user.ID,
		},
		NickName:  user.NickName,
		HeaderImg: user.HeaderImg,
		Phone:     user.Phone,
		Email:     user.Email,
		SideMode:  user.SideMode,
		Enable:    user.Enable,
	})
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// @Param     json {"id":?}
// @Summary   直接重置用户（byid）密码为默认值123456
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.GVA_LOG.Error("重置失败!", zap.Error(err))
		response.FailWithMessage("重置失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("重置成功", c)
}

// @Summary   分页获取用户列表
func (b *BaseApi) GetUserList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoList(pageInfo) //分页获取用户列表，（page，pageSize）->(limit,offset)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 获取用户信息byUUID
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{"userInfo": ReqUser}, "获取成功", c)
}
