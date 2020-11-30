package md5

import (
	"crypto/md5"
	"fmt"
)

/**
加密
*/
func Enc(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

/**
解密
*/
func Dec(str string) string {
	return ""
}
