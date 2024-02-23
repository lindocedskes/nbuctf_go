package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/response"
	"github.com/lindocedskes/model/system"
	systemRes "github.com/lindocedskes/model/system/response"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"strconv"
)

// 大文件分块和计算大文件及其切片的md5在前端完成的，之后拼接也由前端完成
// 前端传来文件片与当前片为什么文件的第几片
// 后端拿到以后比较次分片是否上传 或者是否为不完全片
// 前端发送每片多大
// 前端告知是否为最后一片且是否完成

// @Summary   文件断点续传到服务器，存储路径 breakpointDir/fileMd5/下存储fileName_contentNumber 的切片
func (b *FileUploadAndDownloadApi) BreakpointContinue(c *gin.Context) {
	fileMd5 := c.Request.FormValue("fileMd5")
	fileName := c.Request.FormValue("fileName")
	chunkMd5 := c.Request.FormValue("chunkMd5")                        //该块Md5的值，验证是否发生传输错误
	chunkNumber, _ := strconv.Atoi(c.Request.FormValue("chunkNumber")) //块的编号
	chunkTotal, _ := strconv.Atoi(c.Request.FormValue("chunkTotal"))   //总块数
	_, FileHeader, err := c.Request.FormFile("file")                   //上传的分块文件
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	f, err := FileHeader.Open() //打开文件
	if err != nil {
		global.GVA_LOG.Error("文件读取失败!", zap.Error(err))
		response.FailWithMessage("文件读取失败", c)
		return
	}
	defer func(f multipart.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	cen, _ := io.ReadAll(f)             //读取文件内容存储到cen
	if !utils.CheckMd5(cen, chunkMd5) { //读取到的内容的MD5值是否与预期的块的MD5值匹配
		global.GVA_LOG.Error("检查md5失败!", zap.Error(err))
		response.FailWithMessage("检查md5失败", c)
		return
	}
	//上传文件时检测，FindOrCreate 逻辑上完整的文件记录是否SysFile表存在记录和传输完成——只检查逻辑上完整的文件记录，返回一个逻辑文件记录
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找或创建逻辑大文件记录失败!", zap.Error(err))
		response.FailWithMessage("查找或创建逻辑大文件记录失败", c)
		return
	}
	//路径 breakpointDir/fileMd5/下存储fileName_contentNumber 的切片
	pathC, err := utils.BreakPointContinue(cen, fileName, chunkNumber, chunkTotal, fileMd5)

	if err != nil {
		global.GVA_LOG.Error("断点续传失败!", zap.Error(err))
		response.FailWithMessage("断点续传失败", c)
		return
	}
	//创建文件切片记录，存储位置，关联逻辑完整文件的SysFileID，切片号
	if err = fileUploadAndDownloadService.CreateFileChunk(file.ID, pathC, chunkNumber); err != nil {
		global.GVA_LOG.Error("创建文件切片记录失败!", zap.Error(err))
		response.FailWithMessage("创建文件切片记录失败", c)
		return
	}
	//每次存储后判断是否传输完所有切片
	err = fileUploadAndDownloadService.CheckFileChunksDone(file.ID, chunkTotal)
	response.OkWithMessage("切片创建成功", c)
}

// @Summary   查找文件 by fileMd5 & fileName，返回已传输的每个切片信息
func (b *FileUploadAndDownloadApi) FindFile(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	chunkTotal, _ := strconv.Atoi(c.Query("chunkTotal"))
	file, err := fileUploadAndDownloadService.FindOrCreateFile(fileMd5, fileName, chunkTotal)
	if err != nil {
		global.GVA_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage("查找失败", c)
	} else {
		response.OkWithDetailed(systemRes.FileResponse{File: file}, "查找成功", c)
	}
}

// 整合切片文件 ,存储到 finishDir，fileName命名，返回新文件的路径
func (b *FileUploadAndDownloadApi) BreakpointContinueFinish(c *gin.Context) {
	fileMd5 := c.Query("fileMd5")
	fileName := c.Query("fileName")
	filePath, err := utils.MakeFile(fileName, fileMd5)
	if err != nil {
		global.GVA_LOG.Error("文件创建失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.FilePathResponse{FilePath: filePath}, "文件创建失败", c)
	} else {
		response.OkWithDetailed(systemRes.FilePathResponse{FilePath: filePath}, "文件创建成功", c)
	}
}

// @Summary   删除缓存切片，这时候把逻辑大文件的"IsFinish":  true,
func (b *FileUploadAndDownloadApi) RemoveChunk(c *gin.Context) {
	var file system.SysFile
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.RemoveChunk(file.FileMd5) // 移除本地上的切片
	if err != nil {
		global.GVA_LOG.Error("缓存切片删除失败!", zap.Error(err))
		return
	}
	err = fileUploadAndDownloadService.DeleteFileChunk(file.FileMd5, file.FilePath) //软删除缓存切片记录，并把逻辑大文件的"IsFinish":  true,
	if err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("缓存切片删除成功", c)
}
