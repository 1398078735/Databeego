package blockchain

import (
	"Datarenzheng1010/models"
	"errors"
	"github.com/bolt-master"
	"math/big"
)

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LAST_HASH = "lasthash"

var CHAIN *BlockChain

//区块链结构体的定义，代表的是一条区块链
//功能:1,将新区块数据与已存区块链接
//2,查询某个区块的数据信息
//3,遍历区块信息
type BlockChain struct {
	LastHash []byte   //表示区块链中最细区块的哈希,用于查找最新的区块内容
	BoltDb   *bolt.DB //区块链中操作区块链数据文件的数据库操作对象
}

//创建一个区块链
func NewBlockChain() *BlockChain {
	var bc *BlockChain
	//先打开文件
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)

	//查询chain.db文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME)) //假设有桶
		if bucket == nil {                       //说明没有桶,要创建新桶
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
			if err != nil {
				panic(err.Error())
			}
		}
		//
		lastHash := bucket.Get([]byte(LAST_HASH))
		if len(lastHash) == 0 { //桶中没有lasthash记录，需要新建创世区块并保存
			//创世区块
			genesis := CreateGenesis()
			//区块序列化以后的数据
			genesisBytes := genesis.Serialize()
			//把创世区块存贮到桶中
			//序列化
			bucket.Put(genesis.Hash, genesisBytes)
			//更新最新区块的哈希值
			bucket.Put([]byte(LAST_HASH), genesis.Hash)
			bc = &BlockChain{
				LastHash: genesis.Hash,
				BoltDb:   db,
			}
		} else { //桶中已有lasthash的记录,不再需要创建创世区块,只需要读取即可
			lasthash := bucket.Get([]byte(LAST_HASH))
			bc = &BlockChain{
				LastHash: lasthash,
				BoltDb:   db,
			}
		}
		return nil
	})
	CHAIN = bc
	return bc
}

//方法
//保存数据到区块链中:先生成一个新区块,然后将新区块添加到区块链当中
func (bc *BlockChain) SaveData(data []byte) (Block, error) {
	//从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	//error的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("读取区块链数据失败")
			return err
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})

	//新建一个区块
	NewBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//放入文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//把新创建的区块存入到boltdb数据库中
		bucket.Put(NewBlock.Hash, NewBlock.Serialize())
		//更新LASTHASH对应的值,更新为最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH), NewBlock.Hash)
		bc.LastHash = NewBlock.Hash //将区块链实例的LASTHASH更新为最新区块的哈希
		return nil
	})
	//返回值语句,newblock,err，其中err可能包含错误信息
	return NewBlock, err
}

//该方法用于完成根据用户输入的区块高度查询对应的区块信息
func (bc *BlockChain) QueryBlockByHeight(height int64) (*Block,error) {
	if height < 0{
		return nil,nil
	}
	db := bc.BoltDb
	var errs error
	var eachBlock *Block

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			errs = errors.New("读取区块链数据失败")
			return errs
		}
		//each:每一个
		eachHash := bc.LastHash
		for {
			//最后一个区块的byte类型
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
			eachBlock, errs = DeSerialize(eachBlockBytes)
			if errs != nil {
				return errs
			}
			if eachBlock.Height < height {
				break
			}
			if eachBlock.Height == height { //跳出循环
				break
			}
			//如果高度匹配不满足用户的条件
			eachHash = eachBlock.PrevHash
		}
			return nil
		})
	return eachBlock,errs
}

//该方法用于遍历区块链blockchaini.db文件，并将所有的区块查出并返回
func (bc BlockChain) QueryAllBlocks()([]*Block,error){
	blocks :=make([]*Block,0)//blocks是一个切片容器,用于盛放查询到的区块
	db := bc.BoltDb
	//从文件中查询所有的区块
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket :=tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil{
			err = errors.New("查询区块链数据失败")
		}
		//bucket存在
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)//默认值0的大整数
		for{
			//根据区块的哈希值获取对应的区块
			eachBlockBytes := bucket.Get([]byte(eachHash))
			//反序列化操作
			eachBlock, _ := DeSerialize(eachBlockBytes)
			//将遍历到每一个区块放入到切片容器当中
			blocks = append(blocks,eachBlock)
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0{//找到了创世区块
				break
			}
			//不满足条件,没有找到创世区块
			eachHash = eachBlock.PrevHash
		}
		return nil
	})

	return blocks,err
}

//该方法用于根据用户输入的认证号查询到对应的区块信息
func (bc BlockChain) QueryBlockByCertId(cert_id string)*Block{
	db := bc.BoltDb
	var err error
	var block *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询链上数据失败")
			return err
		}
		eachHash := bc.LastHash
		var eachBlock *Block
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)//默认值0的大整数
		for {//无限循环查询cardid的值
			//最后一个区块的byte类型
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
			eachBlock, err = DeSerialize(eachBlockBytes)
			if err != nil {
				break
			}
			//将遍历到的区块中的数据跟用户提供的认证号进行比较
			record,err:=models.DeSerializeCertRecord(eachBlock.Data)
			if err != nil {
				err = errors.New("查询链上数据失败")
				return err
			}
			if string(record.CertId) == cert_id {//说明找到区块了
				block = eachBlock
				break
			}
			eachBig.SetBytes(eachBlock.PrevHash)
			if eachBig.Cmp(zeroBig) == 0{//找到了创世区块
				break
			}
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return block
}