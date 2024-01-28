package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

//	func (server *_server) InitServer(address string, router *gin.Engine) *http.Server {
//		return &http.Server{
//			Addr:           address,
//			Handler:        router,
//			ReadTimeout:    20 * time.Second,
//			WriteTimeout:   20 * time.Second,
//			MaxHeaderBytes: 1 << 20,
//		}
//	}

// 实现 server 接口的，功能：初始化了一个服务器
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router) // endless.NewServer() 返回一个 *Server 实例，实现了 server 接口
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
