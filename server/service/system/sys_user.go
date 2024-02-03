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
	"time"
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

// @description: 修改一个用户的权限，更新用户的主角色ID
func (userService *UserService) SetUserAuthority(id uint, authorityId uint) (err error) {
	//查询连接表 SysUserAuthority，用户id具有角色id
	assignErr := global.NBUCTF_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityId).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	//更新用户的主角色ID
	err = global.NBUCTF_DB.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// @description: 设置一个用户组的权限，指定用户 :对Sys_User_Authority 表先删除用户id的所有角色，再创建用户id的所有角色记录。并更新第一个角色为主角色
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	//Transaction 方法传入一个回调函数，所有的数据库操作都在同一个事务中执行
	return global.NBUCTF_DB.Transaction(func(tx *gorm.DB) error {
		TxErr := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error //删除用户id的所有角色
		if TxErr != nil {
			return TxErr
		}
		var useAuthority []system.SysUserAuthority
		for _, v := range authorityIds { //遍历角色id，加入SysUserAuthority类型 切片中
			useAuthority = append(useAuthority, system.SysUserAuthority{
				SysUserId: id, SysAuthorityAuthorityId: v,
			})
		}
		TxErr = tx.Create(&useAuthority).Error //创建用户id的所有角色（UserId，AuthorityId）
		if TxErr != nil {
			return TxErr
		}
		//并使用 Update 方法更新用户的 authority_id 字段为 authorityIds 切片的第一个元素
		TxErr = tx.Where("id = ?", id).First(&system.SysUser{}).Update("authority_id", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		// 返回 nil 提交事务
		return nil
	})
}

// @description: 删除用户，按id 删除对应用户
func (userService *UserService) DeleteUser(id int) (err error) {
	return global.NBUCTF_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

// @description: 更新用户信息
func (userService *UserService) SetUserInfo(req system.SysUser) error {
	return global.NBUCTF_DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"header_img": req.HeaderImg,
			"phone":      req.Phone,
			"email":      req.Email,
			"side_mode":  req.SideMode,
			"enable":     req.Enable,
		}).Error
}

// @description: 设置自己的用户信息
func (userService *UserService) SetSelfInfo(req system.SysUser) error {
	return global.NBUCTF_DB.Model(&system.SysUser{}).
		Where("id=?", req.ID).
		Updates(req).Error
}
