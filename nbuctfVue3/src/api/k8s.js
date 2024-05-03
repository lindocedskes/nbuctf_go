import service from '@/utils/request'

// @Summary 开启容器
export const openContainerApi = (data) => {
  return service({
    url: '/k8s/openContainer',
    method: 'POST',
    data
  })
}

// @Summary 查看容器状态
export const checkContainerApi = (data) => {
  return service({
    url: '/k8s/checkContainer',
    method: 'POST',
    data
  })
}
// @Summary 关闭容器
export const closeContainerApi = (data) => {
  return service({
    url: '/k8s/closeContainer',
    method: 'POST',
    data
  })
}

// @Summary 关闭所有容器
export const closeAllContainerApi = () => {
  return service({
    url: '/k8s/closeAllContainer',
    method: 'GET'
  })
}
