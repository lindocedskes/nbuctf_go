package utils

import (
	"math/rand"
	"time"
)

// 默认的字符集
var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~!@#$%^&*()_+-={}[]|;:'<>,.")

// 生成指定长度的随机字符串
func RandomString(length int, allowedChars ...[]rune) string { //rune 是 int32 的别名，用于区分字符值和整数值
	var letters []rune
	if len(allowedChars) == 0 { //未传入 字符集
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	stringContainer := make([]rune, length) //与 new 不同，make(类型,长度) rune类型数组
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range stringContainer {
		stringContainer[i] = letters[r.Intn(len(letters))]
	}

	return string(stringContainer)
}
