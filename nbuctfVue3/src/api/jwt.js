import service from '@/utils/request'
// @Summary jwt加入黑名单
export const jsonInBlacklist = () => {
  return service({
    url: '/jwt/jsonInBlacklist',
    method: 'post'
  })
}
