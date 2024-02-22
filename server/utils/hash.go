package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// md5 哈希，哈希值唯一且固定长度
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)                        //将输入的字节切片str写入到哈希对象h中
	return hex.EncodeToString(h.Sum(b)) //h.Sum(b)计算并返回哈希对象h的哈希值，如果提供了参数b，则哈希值会追加到b的尾部。再转为16进制字符串
}
