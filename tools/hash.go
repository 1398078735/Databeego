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

//对区块数据进行sha256计算
func SHA256HashBlock(bs []byte) []byte{
	//将block结构体数据转换成[]byte类型
	//HeightBytes,_ := Int64ToByte(block.Height)
	//timeStampBytes,_ := Int64ToByte(block.TimeStamp)
	//versionBytes:= StringToBytes(block.Version)
	//
	//var blockBytes []byte
	////bytes.join拼接
	//bytes.Join([][]byte{
	//	HeightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//},[]byte{})
	//将转换后的切片输入write方法
	hash256 := sha256.New()
	hash256.Write(bs)
	hash := hash256.Sum(nil)
	return hash
}

//io输入和输出
func Md5HashReader(reader io.Reader)string{
	hashmd5 := md5.New()
	fileBytes, _ := ioutil.ReadAll(reader)
	hashmd5.Write(fileBytes)
	hashBytes := hashmd5.Sum(nil)
	return hex.EncodeToString(hashBytes)
}