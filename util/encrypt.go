package util

import (
	"crypto/md5"
	"fmt"
)

// md5加密 字符串
func EncryptMyStr(text string) string {
	data := []byte(text)
	has := md5.Sum(data)
	enText := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return enText
}

func EncryptPwd(text string) string {
	data := []byte(text + "company")
	has := md5.Sum(data)
	enText := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return enText
}

func PF(v interface{}) {
	fmt.Printf("v : %+v\n", v)
}
