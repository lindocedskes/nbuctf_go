import service from '@/utils/request'
// @Summary 获取角色列表
export const getAuthorityList = (data) => {
  return service({
    url: '/authoritybyadmin/getAuthorityList',
    method: 'post',
    data
  })
}

// @Summary 删除角色
export const deleteAuthority = (data) => {
  return service({
    url: '/authoritybyadmin/deleteAuthority',
    method: 'post',
    data
  })
}

// @Summary 创建角色
export const createAuthority = (data) => {
  return service({
    url: '/authoritybyadmin/createAuthority',
    method: 'post',
    data
  })
}

// @Tags authority
// @Summary 拷贝角色
export const copyAuthority = (data) => {
  return service({
    url: '/authority/copyAuthority',
    method: 'post',
    data
  })
}

// @Summary 设置角色资源权限
export const setDataAuthority = (data) => {
  return service({
    url: '/authoritybyadmin/setDataAuthority',
    method: 'post',
    data
  })
}

// @Summary 修改角色
export const updateAuthority = (data) => {
  return service({
    url: '/authoritybyadmin/updateAuthority',
    method: 'put',
    data
  })
}
