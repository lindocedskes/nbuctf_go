package system

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	systemReq "github.com/lindocedskes/model/system/request"
	"gorm.io/gorm"
	"strconv"
)

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService) //全局实例，可以在任何需要处理角色相关操作的地方使用
var ErrRoleExistence = errors.New("存在相同角色id")

// @description: 创建一个角色
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {

	if err = global.NBUCTF_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence //存在相同角色id
	}

	e := global.NBUCTF_DB.Transaction(func(tx *gorm.DB) error { //事务
		if err = tx.Create(&auth).Error; err != nil {
			return err
		}

		auth.SysBaseMenus = systemReq.DefaultMenu() //角色的 SysAuthority.SysBaseMenus 字段 ，创建角色时添加默认menu
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := systemReq.DefaultCasbin() //创建角色时添加默认casbin权限 todo 后续调整新建用户的默认路由权限
		authorityId := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos { //将casbinInfos转换为rules
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules) //添加规则
	})

	return auth, e
}
