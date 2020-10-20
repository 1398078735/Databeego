package tools

import (
	"bytes"
	"encoding/binary"
)

//将一个int64转换为[]byte切片
func Int64ToByte(num int64)([]byte,error){
	//buffer:缓冲区
	buff := new(bytes.Buffer)//通过new实例化一个暖冲区
	//buff.Write()//通过write方法向缓冲区写入数据
	//buff.Bytes()通过bytes方法从缓冲区中获取数据
	//binary二进制
	//大端位序排列binary.BigEndian
	//小端位序排列
	err:=binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		return nil,err
	}
	//从缓冲区读取数据
	return buff.Bytes(),nil
}

func StringToBytes(data string)[]byte{
	return []byte(data)
}