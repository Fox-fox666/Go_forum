package common

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "chaojibbt"

func EncryptPassword(s string) string {
	// 将密码和盐值拼接在一起
	data := []byte(s + salt)

	// 计算哈希值
	hash := md5.Sum(data)

	// 将哈希值转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}
