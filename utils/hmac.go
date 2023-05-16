package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
)

func Hmac(key string, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}