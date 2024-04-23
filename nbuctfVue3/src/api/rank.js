import service from '@/utils/request'

// @Summary 分页获取排行榜列表
export const getRankList = (data) => {
  return service({
    url: '/game/rank',
    method: 'post',
    data: data
  })
}
