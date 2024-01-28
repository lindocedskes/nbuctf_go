package system

// service 层的代码，对数据库的操作
import (
	"errors"
	"fmt"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/utils"
)

type UserService struct{}

// @description: 用户登录
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.NBUCTF_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	// 预加载，TODO： 权限（Authority）和角色（Authorities）表，按username查询
	err = global.NBUCTF_DB.Where("username = ?", u.Username).Preload("Authorities").Preload("Authority").First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user) // 获取用户角色默认菜单
	}
	return &user, err
}
