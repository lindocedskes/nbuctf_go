package router

import (
	"github.com/lindocedskes/router/example"
	"github.com/lindocedskes/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup  //系统路由
	Example example.RouterGroup //测试路由
}

var RouterGroupApp = new(RouterGroup)
