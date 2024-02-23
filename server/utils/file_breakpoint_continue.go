package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	breakpointDir = "./uploads/breakpointDir/"
	finishDir     = "./uploads/fileDir/"
)

func CheckMd5(content []byte, chunkMd5 string) (CanUpload bool) {
	fileMd5 := MD5V(content)
	//fmt.Println(fileMd5) //测试
	if fileMd5 == chunkMd5 {
		return true // 可以继续上传
	} else {
		return false // 切片不完整，废弃
	}
}

// @description: 断点续传，路径 breakpointDir/fileMd5/下存储fileName_contentNumber 的切片
func BreakPointContinue(content []byte, fileName string, contentNumber int, contentTotal int, fileMd5 string) (string, error) {
	path := breakpointDir + fileMd5 + "/" //存储路径
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return path, err
	}
	pathC, err := makeFileContent(content, fileName, path, contentNumber) //存储
	return pathC, err
}

// @description: 存储切片内容，路径FileDir下存储fileName_contentNumber 的切片
func makeFileContent(content []byte, fileName string, FileDir string, contentNumber int) (string, error) {
	//防止路径遍历攻击
	if strings.Index(fileName, "..") > -1 || strings.Index(FileDir, "..") > -1 {
		return "", errors.New("文件名或路径不合法")
	}
	path := FileDir + fileName + "_" + strconv.Itoa(contentNumber) //路径+文件命名格式
	f, err := os.Create(path)
	if err != nil {
		return path, err
	} else {
		_, err = f.Write(content) //写入文件内容
		if err != nil {
			return path, err
		}
	}
	defer f.Close()
	return path, nil
}

// @description: 整合切片文件 ,存储到 finishDir，fileName命名，返回新文件的路径和错误
func MakeFile(fileName string, FileMd5 string) (string, error) {
	rd, err := os.ReadDir(breakpointDir + FileMd5) //读取目录（存一个文件的所有切片）的多个文件
	if err != nil {
		return finishDir + fileName, err
	}
	_ = os.MkdirAll(finishDir, os.ModePerm)                                              //创建整合切片文件存储目录
	fd, err := os.OpenFile(finishDir+fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644) //打开或创建一个新的文件，以fileName命名
	if err != nil {
		return finishDir + fileName, err
	}
	defer fd.Close()
	for k := range rd { //整合
		content, _ := os.ReadFile(breakpointDir + FileMd5 + "/" + fileName + "_" + strconv.Itoa(k))
		_, err = fd.Write(content) //将内容写入新文件
		if err != nil {
			_ = os.Remove(finishDir + fileName) //如果在写入过程中发生错误，函数删除新文件
			return finishDir + fileName, err
		}
	}
	return finishDir + fileName, nil //返回新文件的路径和错误
}

// @description: 移除切片
func RemoveChunk(FileMd5 string) error {
	err := os.RemoveAll(breakpointDir + FileMd5)
	return err
}
