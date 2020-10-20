package main

import (
	"Datarenzheng1010/blockchain"
	"Datarenzheng1010/db_mysql"
	_ "Datarenzheng1010/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main(){

	block0 := blockchain.CreateGenesis()
	block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	fmt.Println(block1)
	return

	db_mysql.Connect()
	//静态资源文件路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()

}
