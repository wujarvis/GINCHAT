package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	// 小写
	h := md5.New()
	h.Write([]byte(data))
	tmpStr := h.Sum(nil)
	return hex.EncodeToString(tmpStr)
}

func MD5Encode(data string) string {
	// 大写
	return strings.ToUpper(Md5Encode(data))
}

func MakePassword(plainPwd, salt string) string {
	// 加密
	return Md5Encode(plainPwd + salt)
}

func ValidPassword(clientPassword, salt, serverPassword string) bool {
	// 校验
	return Md5Encode(clientPassword+salt) == serverPassword
}
