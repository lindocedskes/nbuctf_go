package upload

import (
	"errors"
	"github.com/lindocedskes/global"
	"github.com/lindocedskes/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type Local struct{}

// @description: 上传文件
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并md5加密返回hash字符串
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名：哈希字符_当前时间.扩展名
	filename := name + "_" + time.Now().Format("20060102150405") + ext //20060102150405 固定值表示格式显示当前时间
	// 尝试创建此路径，根据设置的存储路径，os.ModePerm表示所有用户都有最大的权限
	mkdirErr := os.MkdirAll(global.NBUCTF_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.NBUCTF_CONFIG.Local.StorePath + "/" + filename
	filepath := global.NBUCTF_CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p) //创建一个新的文件，由out指向
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() failed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

// @description: 删除文件 by文件名
func (*Local) DeleteFile(key string) error {
	p := global.NBUCTF_CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.NBUCTF_CONFIG.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
