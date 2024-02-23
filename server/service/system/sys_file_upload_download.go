package system

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/model/common/request"
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

// @description: 分页获取文件列表，+支持对文件name模糊查询-通过%keyword%
func (e *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.NBUCTF_DB.Model(&system.SysFileUploadAndDownload{})
	var fileLists []system.SysFileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%") //对文件name 模糊查询
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

// @description: 按文件id查找文件记录
func (e *FileUploadAndDownloadService) FindFile(id uint) (system.SysFileUploadAndDownload, error) {
	var file system.SysFileUploadAndDownload
	err := global.NBUCTF_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// @description: 删除文件记录
func (e *FileUploadAndDownloadService) DeleteFile(file system.SysFileUploadAndDownload) (err error) {
	var fileFromDb system.SysFileUploadAndDownload
	fileFromDb, err = e.FindFile(file.ID) //按文件id查找文件记录
	if err != nil {
		return
	}
	oss := upload.NewOss()                                //oss对象
	if err = oss.DeleteFile(fileFromDb.Key); err != nil { //按key（实际存储文件标识名）执行删除存储在oss上对应文件（）
		return errors.New("文件删除失败")
	}
	err = global.NBUCTF_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error //数据库中删除改记录
	return err
}

// 编辑文件名namebyid
func (e *FileUploadAndDownloadService) EditFileName(file system.SysFileUploadAndDownload) (err error) {
	var fileFromDb system.SysFileUploadAndDownload
	return global.NBUCTF_DB.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}
