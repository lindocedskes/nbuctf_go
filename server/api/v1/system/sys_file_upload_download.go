package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/request"
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

// @Summary   分页获取文件列表，+支持对文件name模糊查询-通过%keyword%
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo) //分页 + keyword 模糊查询：%keyword%
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// @Summary   删除文件by id
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file system.SysFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// 编辑文件名namebyid（实际存储文件名为key）
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file system.SysFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.EditFileName(file)
	if err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}
