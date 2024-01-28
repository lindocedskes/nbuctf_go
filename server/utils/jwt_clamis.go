package utils

import (
	"github.com/gin-gonic/gin"
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
