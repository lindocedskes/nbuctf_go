package system

import "errors"

const (
// Mysql           = "mysql"
// Pgsql           = "pgsql"
// InitSuccess     = "\n[%v] --> 初始数据成功!\n"
// InitDataExist   = "\n[%v] --> %v 的初始数据已存在!\n"
// InitDataFailed  = "\n[%v] --> %v 初始数据失败! \nerr: %+v\n"
// InitDataSuccess = "\n[%v] --> %v 初始数据成功!\n"
)

const (
// InitOrderSystem   = 10
// InitOrderInternal = 1000
// InitOrderExternal = 100000
)

var (
	ErrMissingDBContext        = errors.New("missing db in context")
	ErrMissingDependentContext = errors.New("missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("db type mismatch")
)
