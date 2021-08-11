package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)

}
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func MakePasswd(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

func ValidatePasswd(plainpwd, salt, passwd string) bool {
	return Md5Encode(plainpwd+salt) == passwd
}
