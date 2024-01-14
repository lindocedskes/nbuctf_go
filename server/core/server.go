package core

import (
	"fmt"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/initialize"
	"go.uber.org/zap"
	"time"
)

type server interface { // 定义服务接口
	ListenAndServe() error
}

func RunServer() {
	// todo JWT,redis

	Router := initialize.Routers() // todo 初始化路由,调用api，resonses

	address := fmt.Sprintf(":%d", global.NBUCTF_CONFIG.System.Port)
	s := initServer(address, Router)  // 初始化服务，返回服务实例。server_other.go/initServer()实现了server接口
	time.Sleep(10 * time.Microsecond) // 保证文本顺序输出
	global.GVA_LOG.Info("server run success on ", zap.String("port", address))

	// 启动服务
	global.GVA_LOG.Error(s.ListenAndServe().Error()) // 启动服务，通过调用ListenAndServe()方法实现

}
