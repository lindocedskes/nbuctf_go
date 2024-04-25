import service from '@/utils/request'

// @Summary 获取公告
export const GetAnnouncements = () => {
  return service({
    url: '/announcement/list',
    method: 'get'
  })
}

// @Summary 添加公告
export const AddAnnouncement = (data) => {
  return service({
    url: '/announcementadmin/create',
    method: 'post',
    data
  })
}
// @Summary 删除公告
export const DeleteAnnouncement = (data) => {
  return service({
    url: '/announcementadmin/delete',
    method: 'post',
    data
  })
}

// @Summary 修改公告
export const UpdateAnnouncement = (data) => {
  return service({
    url: '/announcementadmin/update',
    method: 'post',
    data
  })
}
