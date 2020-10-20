package blockchain

import (
	"Datarenzheng1010/tools"
	"bytes"
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

	//将block结构体数据转换成[]byte类型
	HeightBytes,_ := tools.Int64ToByte(block.Height)
	timeStampBytes,_ := tools.Int64ToByte(block.TimeStamp)
	versionBytes:= tools.StringToBytes(block.Version)
	var blockBytes []byte
	//bytes.join拼接
	bytes.Join([][]byte{
		HeightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	},[]byte{})
	//调用hash计算,对区块进行sha256计算
	block.Hash = tools.SHA256HashBlock(blockBytes)
	return block
}


//创建创世区块
func CreateGenesis() Block{
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return genesisBlock
}
