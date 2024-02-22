package system

import (
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/system"
	"github.com/lindocedskes/utils/upload"
	"mime/multipart"
	"strings"
)

type FileUploadAndDownloadService struct{} //service 层 文件上传下载调用结构体

// @description: 根据配置文件判断是文件上传到本地或者七牛云。 *multipart.FileHeader 一个MIME多部分文件的头部信息
func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (file system.SysFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header) //上传文件，header，返回存储路径
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := system.SysFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	if noSave == "0" {
		return f, e.UploadRecord(f) //0 为记录标记，创建文件上传记录
	}
	return f, nil
}

// @description: 创建文件上传记录
func (e *FileUploadAndDownloadService) UploadRecord(file system.SysFileUploadAndDownload) error {
	return global.NBUCTF_DB.Create(&file).Error
}
