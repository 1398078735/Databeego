package tools

import (
	"io"
	"os"
)

//tools工具的意思:用于把重复的代码进行封装减少代码的重复
//保存文件方法
func SaveFile(fileName string,file io.Reader)(int64,error){
	saveFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		return -1,err
	}
	length , err := io.Copy(saveFile, file)
	if err != nil {
		return -1,err
	}
	return length,nil
}
