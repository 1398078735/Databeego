package main

import (
	"Datarenzheng1010/blockchain"
	"Datarenzheng1010/db_mysql"
	_ "Datarenzheng1010/routers"
	"github.com/astaxie/beego"
)


func main(){

	//先准备一条区块链
	blockchain.CHAIN = blockchain.NewBlockChain()

	db_mysql.Connect()
	//静态资源文件路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()

}


//bc := blockchain.NewBlockChain() // 封装
//bc.SaveData([]byte("区块链学院"))
//blocks,err:=bc.QueryAllBlocks()
//if err != nil {
//	fmt.Println(err.Error())
//	return
//}
//
////blocks是一个切片
//for _, block := range blocks{
//	fmt.Printf("区块高度:%d,区块内数据:%s\n",block.Height,string(block.Data))
//}
//
//
//return

//block0 := blockchain.CreateGenesis()
//block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
//fmt.Printf("block1哈希值%x\n",block1.Hash)
//fmt.Printf("block1的上一个哈希%x\n",block1.PrevHash)
//fmt.Printf("block0哈希值%x\n",block0.Hash)

//序列化，将数据从内存当中形式转换为可以持久化存贮在硬盘或者网络传输的形式,称为序列化
//反序列化,将数据从文件中或者网络中读取,然后转化到计算机内存的过程称为反序列化
//blockjson,_:=json.Marshal(block0)
//fmt.Println("序列化以后的blockjson",string(blockjson))
//block0Bytes := block0.Serialize()
//fmt.Println("序列化以后",block0Bytes)
//deBlock0,err := blockchain.DeSerialize(block0Bytes)
//if err != nil {
//	fmt.Println(err.Error())
//	return
//}
//fmt.Println("反序列化后区块的高度",deBlock0.Height)
//fmt.Printf("反序列化的区块的哈希%x\n",deBlock0.Hash)
//return

//blockN,err := bc.SaveData([]byte("区块链学院"))
////fmt.Printf("最新区块的哈希值%x\n",bc.LastHash)
//if err != nil {
//	fmt.Println(err.Error())
//	return
//}
//fmt.Printf("区块的高度%d\n",blockN.Height)
//fmt.Printf("区块的哈希%x\n",blockN.Hash)
//fmt.Printf("区块的前一个哈希%x\n",blockN.PrevHash)
//block2,err:=bc.QueryBlockByHeight(11)
//if err!=nil {
//	fmt.Println(err.Error())
//
//}
//fmt.Printf("区块的高度%d\n",block2.Height)
//fmt.Println("区块的数据是",string(block2.Data))