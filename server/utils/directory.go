package utils

import (
	"github.com/lindocedskes/global"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
)

// 文件目录是否存在
func PathExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path) // os.Stat获取文件信息
	if err == nil {
		if fileInfo.IsDir() {
			return true, nil
		} else {
			return false, errors.New("check path is not dir")
		}
	}
	if os.IsNotExist(err) { // os.IsNotExist 判断文件是否存在
		return false, nil
	}
	return false, err
}

// 批量创建文件夹
func CreateDir(dirs ...string) error {
	for _, dir := range dirs {
		ok, err := PathExists(dir)
		if err != nil {
			return err
		}
		if !ok {
			global.GVA_LOG.Debug("create directory" + dir)
			err = os.MkdirAll(dir, os.ModePerm) // os.ModePerm：0777
			if err != nil {
				global.GVA_LOG.Error("create directory"+dir, zap.Any(" error:", err)) //写入日志概述+日志字段详情
				return err
			}
		}

	}
	return nil
}
