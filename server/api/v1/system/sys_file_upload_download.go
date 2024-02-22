package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	systemRes "github.com/lindocedskes/model/system/response"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

// @Summary 文件上传
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file system.SysFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")      // 0表示上传文件相关信息记录在表中，否则不记录
	_, header, err := c.Request.FormFile("file") //从请求获取文件，返回：1.multipart.File，提供了对上传的文件数据进行读取的方法。 2.*multipart.FileHeader 包含了文件的元数据、头部信息，并有open() 方法可以读数据
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysFileResponse{File: file}, "上传成功", c)
}
