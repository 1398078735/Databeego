package blockchain

import (
	"Datarenzheng1010/tools"
	"bytes"
	"math/big"
)

const DIFFICULTY  = 10
//工作量证明算法结构体
type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block Block //要找的nonce值对应的区块
}

//实例化一个pow算法的实例
func NewPow(block Block) ProofOfWork{
	t := big.NewInt(1)
	t = t.Lsh(t,255- DIFFICULTY)
	pow := ProofOfWork{
		Target: t,
		Block:  block,
	}
	return pow
}

//run用于寻找合适的nonce值
func (p ProofOfWork) Run() ([]byte,int64){
	var blockHash []byte
	var nonce int64
	nonce = 0

	for {
		block := p.Block
		HeightBytes, _ := tools.Int64ToByte(block.Height)
		timeStampBytes, _ := tools.Int64ToByte(block.TimeStamp)
		versionBytes := tools.StringToBytes(block.Version)
		//bytes.join拼接
		nonceBytes, _ := tools.Int64ToByte(nonce)
		blockBytes := bytes.Join([][]byte{
			HeightBytes,
			timeStampBytes,
			block.PrevHash,
			block.Data,
			versionBytes,
			nonceBytes,
		}, []byte{})

		//区块和尝试的nonce拼接后的hash值
		blockHash = tools.SHA256HashBlock(blockBytes)
		var hashBig *big.Int
		hashBig = new(big.Int) //分配内存空间,为变量分配地址
		hashBig = hashBig.SetBytes(blockHash)
		//fmt.Println("当前的nonce值是:",nonce)

		target := p.Target //目标值
		if hashBig.Cmp(target) == -1 {
			//满足条件停止寻找
			break
		}
		nonce++//自增
	}
	//把符合规则的nonce返回
	return blockHash,nonce
}