package v1

//api/v1/enter.go 区分2个api组
import (
	"github.com/lindocedskes/api/v1/example"
	"github.com/lindocedskes/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
