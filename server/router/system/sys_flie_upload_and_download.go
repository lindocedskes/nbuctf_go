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
		fileUploadAndDownloadRouter.POST("upload", fileUploadAndDownloadApi.UploadFile)         // 上传文件
		fileUploadAndDownloadRouter.POST("getFileList", fileUploadAndDownloadApi.GetFileList)   // 分页获取文件列表，+支持对文件name模糊查询-通过%keyword%
		fileUploadAndDownloadRouter.DELETE("deleteFile", fileUploadAndDownloadApi.DeleteFile)   // 删除指定文件byid
		fileUploadAndDownloadRouter.POST("editFileName", fileUploadAndDownloadApi.EditFileName) // 编辑文件名namebyid（实际存储文件名为key不变）
		//切片上传到oss后再整合,4个调用顺序由前端决定来实现完整的切片上传再合并过程
		fileUploadAndDownloadRouter.POST("breakpointContinue", fileUploadAndDownloadApi.BreakpointContinue)            // 断点续传
		fileUploadAndDownloadRouter.GET("findFile", fileUploadAndDownloadApi.FindFile)                                 // 查询传输成功的切片by fileMd5 & fileName
		fileUploadAndDownloadRouter.POST("removeChunk", fileUploadAndDownloadApi.RemoveChunk)                          // 删除缓存切片，这时候把"IsFinish":  true,
		fileUploadAndDownloadRouter.GET("breakpointContinueFinish", fileUploadAndDownloadApi.BreakpointContinueFinish) // 整合切片文件 ,存储到 finishDir，fileName命名，返回新文件的路径
	}
}
