package v1

import (
	"github.com/lindocedskes/api/v1/example"
	"github.com/lindocedskes/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
