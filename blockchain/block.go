package blockchain

import (
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
	//block.Hash =
	return block
}

//创建创世区块
func CreateGenesis() Block{
	genesisBlock := NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return genesisBlock
}

func CalculateHash(block Block) string {
	//height := tools.Int64ToBytes(block.Height)
	//timestamp := tools.Int64ToBytes(block.TimeStamp)
	//record :=  height + timestamp + block.Data + block.PrevHash
	//h := sha256.New()
	//h.Write([]byte(record))
	//hashed := h.Sum(nil)
	//return hex.EncodeToString(hashed)
}

