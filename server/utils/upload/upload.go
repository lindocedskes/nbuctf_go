package upload

import (
	"github.com/lindocedskes/global"
	"mime/multipart"
)

// OSS 对象存储接口

type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法

func NewOss() OSS {
	switch global.NBUCTF_CONFIG.System.OssType {
	case "local":
		return &Local{} //Local{} 实现了OSS接口
	default:
		return &Local{}
	}
}
