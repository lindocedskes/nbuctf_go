package core

import (
	"fmt"
	"github.com/lindocedskes/core/internal"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.NBUCTF_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.NBUCTF_CONFIG.Zap.Director)
		_ = os.Mkdir(global.NBUCTF_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()        // 获取zapcore.Core[]
	logger = zap.New(zapcore.NewTee(cores...)) // 初始化zap日志库

	if global.NBUCTF_CONFIG.Zap.ShowLine { // 是否显示行号
		logger = logger.WithOptions(zap.AddCaller()) // 开启开发模式，堆栈跟踪
	}
	return logger
}
