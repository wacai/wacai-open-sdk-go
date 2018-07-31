package tools

import(
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"crypto/md5"
	"strings"
)

func HMAC256(message string,secretKey string) string{
	key := []byte(secretKey)
	hmac256Ctx := hmac.New(sha256.New, key)
	hmac256Ctx.Write([]byte(message))
	cipher := hmac256Ctx.Sum(nil)
	
	var hmac256Str = base64.StdEncoding.EncodeToString(cipher)
	hmac256Str = strings.Replace(hmac256Str,"+","-",-1)
	hmac256Str = strings.Replace(hmac256Str,"/","_",-1)
	hmac256Str = strings.Replace(hmac256Str,"=","",-1)

	return hmac256Str
}

func MD5(message string) string{
	md5Ctx := md5.New() 
	md5Ctx.Write([]byte(message)) 
	cipher := md5Ctx.Sum(nil)
	md5Str := base64.StdEncoding.EncodeToString(cipher)
    return md5Str
}