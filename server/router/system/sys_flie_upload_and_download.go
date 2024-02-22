package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lindocedskes/api/v1"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("file")
	fileUploadAndDownloadApi := v1.ApiGroupApp.SystemApiGroup.FileUploadAndDownloadApi
	{
		fileUploadAndDownloadRouter.POST("upload", fileUploadAndDownloadApi.UploadFile) // 上传文件
		//fileUploadAndDownloadRouter.POST("getFileList", fileUploadAndDownloadApi.GetFileList)                           // 获取上传文件列表
		//fileUploadAndDownloadRouter.POST("deleteFile", fileUploadAndDownloadApi.DeleteFile)                             // 删除指定文件
		//fileUploadAndDownloadRouter.POST("editFileName", fileUploadAndDownloadApi.EditFileName)                         // 编辑文件名或者备注
		//fileUploadAndDownloadRouter.POST("breakpointContinue", fileUploadAndDownloadApi.BreakpointContinue)             // 断点续传
		//fileUploadAndDownloadRouter.GET("findFile", fileUploadAndDownloadApi.FindFile)                                  // 查询当前文件成功的切片
		//fileUploadAndDownloadRouter.POST("breakpointContinueFinish", fileUploadAndDownloadApi.BreakpointContinueFinish) // 切片传输完成
		//fileUploadAndDownloadRouter.POST("removeChunk", fileUploadAndDownloadApi.RemoveChunk)                           // 删除切片
	}
}
