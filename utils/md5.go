package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(plainpwd string, salt string) string {
	return MD5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(plainpwd string, salt string, password string) bool {
	return MD5Encode(plainpwd+salt) == password
}
