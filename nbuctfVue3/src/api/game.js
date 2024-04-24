import service from '@/utils/request'

// @Summary 用户获取赛题列表,未隐藏的
export const getGameList = () => {
  return service({
    url: '/game/list',
    method: 'get'
  })
}

//提交flag
export const submitFlag = (data) => {
  return service({
    url: '/game/submitflag',
    method: 'post',
    data
  })
}
