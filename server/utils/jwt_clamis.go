package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	systemReq "github.com/lindocedskes/model/system/request"
	"net"
)

// 设cookie中 "x-token" 字段值是 JWT 的字符串，maxAge是cookie的有效时间
func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host) // 从请求头ip+端口，分割获得host
	if err != nil {                                   //不含端口则报错
		host = c.Request.Host
	}
	//   有效时间，路径，域名，是否安全，是否httponly
	c.SetCookie("x-token", token, maxAge, "/", host, false, false)
}

// 获取cookie或请求Header中 "x-token" 字段值
func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token") //cookie没有，从请求的 Header 中获取
	}
	return token
}

// 清除cookie中 "x-token" 字段值
func ClearToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie("x-token", "", -1, "/", host, false, false)
}

// 从请求的上下文jwt字符串解析还原为Claims
func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID，避免越权
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.ID
	}
}
