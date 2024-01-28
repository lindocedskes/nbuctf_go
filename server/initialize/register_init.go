package initialize

import (
	_ "github.com/lindocedskes/source/system"
)

// 导入需要初始化包，会自动执行包的init()函数
// main.go 已导入initialize包，这里导入source/system包，会自动执行source/system包的init()函数
func init() {
	// do nothing,only import source package so that inits can be registered
}
