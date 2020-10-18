package tools

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

//对一个字符串数据进行md5哈希计算
func Md5HashString(data string)string{
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	bytes := md5Hash.Sum(nil)
	return hex.EncodeToString(bytes)
}

//读取io流中的数据,并对数据进行hash计算,返回sha256 hash值
func SHA256HashReader(reader io.Reader)string{
	hash256 := sha256.New()
	fileBytes, _ := ioutil.ReadAll(reader)
	hash256.Write(fileBytes)
	hashBytes := hash256.Sum(nil)
	return hex.EncodeToString(hashBytes)
}

//io输入和输出
func Md5HashReader(reader io.Reader)string{
	hashmd5 := md5.New()
	fileBytes, _ := ioutil.ReadAll(reader)
	hashmd5.Write(fileBytes)
	hashBytes := hashmd5.Sum(nil)
	return hex.EncodeToString(hashBytes)
}