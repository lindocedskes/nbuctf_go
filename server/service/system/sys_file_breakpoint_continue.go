package system

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// @description: 上传文件时检测，FindOrCreate 逻辑上完整的文件记录是否SysFile表存在记录和传输完成——只检查逻辑上完整的文件记录
func (e *FileUploadAndDownloadService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file system.SysFile, err error) {
	var cfile system.SysFile //逻辑上一个完整的文件信息
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal
	//Find：尝试在数据库中查找具有指定MD5值并且已经完成传输（is_finish = true）
	if errors.Is(global.NBUCTF_DB.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).Preload("SysFileChunk").First(&file).Error, gorm.ErrRecordNotFound) {
		//Create：没有找到指定MD5值并且已经完成传输 -> 查找具有指定MD5值和文件名的文件记录,找到返回该传输未完成的文件记录，未找到则根据 cfile 创建一个新的文件FirstOrCreate(&file, cfile)
		err = global.NBUCTF_DB.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("SysFileChunk").FirstOrCreate(&file, cfile).Error
		//无论是找到还是创建，都会将结果赋值给 file 并直接返回查询到的file
		return file, err
	}
	//Find：找到了具有指定MD5值且完成传输 ，直接返回信息
	return file, err
}

// @description: 创建文件切片记录，存储位置，关联逻辑完整文件的SysFileID，切片号
func (e *FileUploadAndDownloadService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk system.SysFileChunk
	//查询该切片是否已经存在
	err := global.NBUCTF_DB.Where("sys_file_id = ? AND file_chunk_number = ?", id, fileChunkNumber).First(&chunk).Error
	//已经存在,则不处理
	if err == gorm.ErrRecordNotFound {
		chunk.FileChunkPath = fileChunkPath
		chunk.SysFileID = id
		chunk.FileChunkNumber = fileChunkNumber
		err = global.NBUCTF_DB.Create(&chunk).Error
	}

	return err
}

// 判断是否传输完所有切片，更新is_finish完成状态
func (e *FileUploadAndDownloadService) CheckFileChunksDone(id uint, chunkTotal int) error {
	var count int64
	var file system.SysFile
	//查询该切片 的总个数
	err := global.NBUCTF_DB.Model(&system.SysFileChunk{}).Where("sys_file_id = ?", id).Count(&count).Error
	if err != nil {
		global.GVA_LOG.Error("获取文件切片总数失败!", zap.Error(err))
		return err
	}
	if count == int64(chunkTotal) { //切片已上传总数与切片总数相同，更新is_finish完成状态
		return global.NBUCTF_DB.Model(&system.SysFile{}).Where("id = ?", id).First(&file).Update("is_finish", true).Error
	}
	return nil
}
