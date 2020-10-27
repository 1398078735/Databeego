package blockchain

import (
	"bytes"
	"encoding/gob"
	"time"
)

//定义区块结构体,用于表示区块
type Block struct {
	Height    int64  // 表示区块的高度,第几个区块
	TimeStamp int64  //时间戳
	PrevHash  []byte //上一个哈希值
	Data      []byte //数据字段
	Hash      []byte //当前区块的哈希值
	Version   string // 版本号
	Nonce     int64 ///区块对应的nonce值
}

//创建一个新区块
func NewBlock(height int64, preHash []byte, data []byte) Block{
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  preHash,
		Data:      data,
		Version:   "1.0",
	}

	//找nonce值,通过工作量证明算法计算寻找
	pow := NewPow(block)
	hash,nonce := pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	//将block结构体数据转换成[]byte类型
	//HeightBytes,_ := tools.Int64ToByte(block.Height)
	//timeStampBytes,_ := tools.Int64ToByte(block.TimeStamp)
	//versionBytes:= tools.StringToBytes(block.Version)
	//nonceBytes,_:=tools.Int64ToByte(block.Nonce)
	//var blockBytes []byte
	////bytes.join拼接
	//bytes.Join([][]byte{
	//	HeightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//	nonceBytes,
	//},[]byte{})
	////调用hash计算,对区块进行sha256计算
	//block.Hash = tools.SHA256HashBlock(blockBytes)

	//挖矿竞争，获得记账权
	//区块A+n<系统B（bigint)
	return block
}


//创建创世区块
func CreateGenesis() Block{
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return genesisBlock
}

//对区块进行序列化
func (b Block) Serialize() ([]byte){
	buff := new(bytes.Buffer)//缓冲区
	encoder := gob.NewEncoder(buff)
	encoder.Encode(b) // 将区块b放入到序列化编码器当中
	return buff.Bytes()
}

//对区块进行反序列化操作
func DeSerialize(data []byte) (*Block,error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block) // 将区块b放入到序列化编码器当中
	if err != nil {
		return nil,err
	}
	return &block,nil
}

