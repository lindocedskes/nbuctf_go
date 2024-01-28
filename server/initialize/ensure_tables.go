package initialize

//确保在应用程序启动时，所有需要的数据库表都已经创建。todo 不必须
//如果表已经存在，则不会执行任何操作。
import (
	"github.com/lindocedskes/service/system"
)

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// auto run
func init() {
	//system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}
