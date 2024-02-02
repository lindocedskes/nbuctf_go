package system

// service 层的代码，对数据库的操作
import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/utils"
	"gorm.io/gorm"
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

// @description: 用户注册,增加记录
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.NBUCTF_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.NBUCTF_DB.Create(&u).Error
	return u, err
}

// @description: 修改用户密码
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.NBUCTF_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.NBUCTF_DB.Save(&user).Error // 更新密码
	return &user, err
}
