import service from '@/utils/request'
// @Summary 用户登录 获取动态路由
// @Param 可以什么都不填 调一下即可
export const asyncMenu = () => {
  return service({
    url: '/menu/getMenu',
    method: 'post'
  })
}
