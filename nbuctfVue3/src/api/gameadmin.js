import service from '@/utils/request'

// @Summary 获取赛题列表
export const getGameList = () => {
  return service({
    url: '/gameadmin/list',
    method: 'get'
  })
}
// @Summary 删除赛题
export const deleteGame = (data) => {
  return service({
    url: '/gameadmin/deletequestion',
    method: 'post',
    data
  })
}

// @Summary 添加赛题
export const addGame = (data) => {
  return service({
    url: '/gameadmin/createquestion',
    method: 'post',
    data
  })
}

// @Summary 编辑赛题
export const editGame = (data) => {
  return service({
    url: '/gameadmin/editquestion',
    method: 'post',
    data
  })
}
// 添加附件
export const addFile = (data) => {
  return service({
    url: '/gameadmin/addfile',
    method: 'post',
    data
  })
}

// 删除附件 by 题目id
export const deleteFile = (data) => {
  return service({
    url: '/gameadmin/deletefiles',
    method: 'post',
    data
  })
}
