import service from '@/utils/request'

// @Summary 获取用户信息
export const getUserInfo = () => {
  return service({
    url: '/user/getUserInfo',
    method: 'get'
  })
}
// @Summary 设置用户信息
export const setSelfInfo = (data) => {
  return service({
    url: '/user/setSelfInfo',
    method: 'put',
    data: data
  })
}

// @Summary 获取验证码
export const captcha = (data) => {
  return service({
    url: '/base/captcha',
    method: 'post',
    data: data
  })
}

// @Summary 用户登录
export const login = (data) => {
  return service({
    url: '/base/login',
    method: 'post',
    data: data
  })
}

// @Summary 修改密码
export const changePassword = (data) => {
  return service({
    url: '/user/changePassword',
    method: 'post',
    data: data
  })
}

// 设置用户信息
export const setUserInfo = (data) => {
  return service({
    url: '/user/setUserInfo',
    method: 'put',
    data: data
  })
}

// @Summary 设置用户权限
export const setUserAuthority = (data) => {
  return service({
    url: '/user/setUserAuthority',
    method: 'post',
    data: data
  })
}
// @Summary 删除用户
export const deleteUser = (data) => {
  return service({
    url: '/user/deleteUser',
    method: 'delete',
    data: data
  })
}
// @Summary 设置用户权限
export const setUserAuthorities = (data) => {
  return service({
    url: '/user/setUserAuthorities',
    method: 'post',
    data: data
  })
}
// @Summary 分页获取用户列表
export const getUserList = (data) => {
  return service({
    url: '/user/getUserList',
    method: 'post',
    data: data
  })
}
// @Summary 获取用户信息
export const resetPassword = (data) => {
  return service({
    url: '/user/resetPassword',
    method: 'post',
    data: data
  })
}
// @Summary 用户注册
export const register = (data) => {
  return service({
    url: '/user/registerbyadmin',
    method: 'post',
    data: data
  })
}
