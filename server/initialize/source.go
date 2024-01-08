package initialize

import (
	"context"
	"fmt"
	"github.com/lindocedskes/common"
	"github.com/lindocedskes/source"
)

// 表数据初始化 Casbin、Role 和 User
func InitSource() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "db", common.NBUCTF_DB)

	if registerCasbin(ctx) {
		common.NBUCTF_CASBIN = source.GetCasbin()
	}
	if registerRole(ctx) {
		fmt.Println("role table data init success")
	}
	if registerUser(ctx) {
		fmt.Println("user table data init success")
	}
}

// 注册 Casbin
func registerCasbin(ctx context.Context) bool {
	initCasbin := source.InitCasbin{}
	if initCasbin.TableCreated(ctx) {
		if initCasbin.DataInserted(ctx) {
			return true
		}
		_, err := initCasbin.InitData(ctx)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}

// 注册 Role
func registerRole(ctx context.Context) bool {
	initRole := source.InitRole{}
	if initRole.TableCreated(ctx) { //判断是否存在表
		if initRole.DataInserted(ctx) { //判断是否存在数据
			return true
		}
		_, err := initRole.InitData(ctx) //初始化数据
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}

func registerUser(ctx context.Context) bool {
	initUser := source.InitUser{}
	if initUser.TableCreated(ctx) {
		if initUser.DataInserted(ctx) {
			return true
		}
		_, err := initUser.InitData(ctx)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}
