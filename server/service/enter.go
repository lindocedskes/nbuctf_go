package service

import (
	"github.com/lindocedskes/service/example"
	"github.com/lindocedskes/service/system"
)

type ServiceGroup struct { //服务组
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
