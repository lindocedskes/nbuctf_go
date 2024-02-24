package system

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/model/system/request"
	"github.com/lindocedskes/model/system/response"
	"gorm.io/gorm"
)

type AuthorityBtnService struct{}

// 按角色id & 菜单id 查询菜单按钮id，返回列表包含所有的菜单按钮id
func (a *AuthorityBtnService) GetAuthorityBtn(req request.SysAuthorityBtnReq) (res response.SysAuthorityBtnRes, err error) {
	var authorityBtn []system.SysAuthorityBtn
	err = global.NBUCTF_DB.Find(&authorityBtn, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authorityBtn {
		selected = append(selected, v.SysBaseMenuBtnID)
	}
	res.Selected = selected
	return res, err
}

// 按角色id & 菜单id 删除已有菜单按钮id，再重新创建（AuthorityId，SysMenuID，SysBaseMenuBtnID）
func (a *AuthorityBtnService) SetAuthorityBtn(req request.SysAuthorityBtnReq) (err error) {
	return global.NBUCTF_DB.Transaction(func(tx *gorm.DB) error {
		var authorityBtn []system.SysAuthorityBtn
		err = tx.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ? and sys_menu_id = ?", req.AuthorityId, req.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range req.Selected {
			authorityBtn = append(authorityBtn, system.SysAuthorityBtn{
				AuthorityId:      req.AuthorityId,
				SysMenuID:        req.MenuID,
				SysBaseMenuBtnID: v,
			})
		}
		if len(authorityBtn) > 0 {
			err = tx.Create(&authorityBtn).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

// 查询是否能删除 菜单按钮id
func (a *AuthorityBtnService) CanRemoveAuthorityBtn(ID string) (err error) {
	fErr := global.NBUCTF_DB.First(&system.SysAuthorityBtn{}, "sys_base_menu_btn_id = ?", ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	//作为外键被关联
	return errors.New("此按钮正在被使用无法删除")
}
