package utils

import (
	"strconv"
	"strings"
	"time"
)

// 解析时间 7d ->
func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)         //去除字符串首尾的空白字符
	dr, err := time.ParseDuration(d) //解析字符串d，返回一个Duration类型的值
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") { //判断字符串d中是否包含子串"d"
		index := strings.Index(d, "d") //返回子串"d"在字符串d中第一次出现的位置

		hour, _ := strconv.Atoi(d[:index])          //将字符串d[:index] “d”之前的字符转换为int类型
		dr = time.Hour * 24 * time.Duration(hour)   //将hour转换为Duration类型
		ndr, err := time.ParseDuration(d[index+1:]) //将字符串d[index+1:] “d”之后的字符转换为Duration类型
		if err != nil {
			return dr, nil //7d -> 7*24h
		}
		return dr + ndr, nil //7d 1h -> 7*24h + 1h
	}

	dv, err := strconv.ParseInt(d, 10, 64) //将字符串d转换为int64类型
	return time.Duration(dv), err
}
