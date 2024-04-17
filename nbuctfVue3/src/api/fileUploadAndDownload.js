import service from '@/utils/request'
// @Tags FileUploadAndDownload
// @Summary 分页文件列表
export const getFileList = (data) => {
  return service({
    url: '/fileUploadAndDownload/getFileList',
    method: 'post',
    data
  })
}

// @Tags FileUploadAndDownload
// @Summary 删除文件
export const deleteFile = (data) => {
  return service({
    url: '/fileUploadAndDownload/deleteFile',
    method: 'post',
    data
  })
}

/**
 * 编辑文件名或者备注
 * @param data
 * @returns {*}
 */
export const editFileName = (data) => {
  return service({
    url: '/fileUploadAndDownload/editFileName',
    method: 'post',
    data
  })
}
