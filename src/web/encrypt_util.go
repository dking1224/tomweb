package web

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

func MD5String(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func MD5(data []byte) string {
	md5 := md5.New()
	md5.Write(data)
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}

func UperStringMD5(data string) {
	strings.ToUpper(MD5String(data))
}

func UpperMD5(data []byte) {
	strings.ToUpper(MD5(data))
}

func UrlEncodeBase64(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func UrlDecodeBase64(data string) string {
	str, _ := base64.URLEncoding.DecodeString(data)
	return string(str)
}

func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}

func Sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
