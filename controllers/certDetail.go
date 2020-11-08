package controllers

import (
	"Datarenzheng1010/blockchain"
	"Datarenzheng1010/models"
	"Datarenzheng1010/tools"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CerDetail struct {
	beego.Controller
}

//用于处理浏览器的get请求
func(c *CerDetail) Get(){
	//解析和接受前端页面传递的数据cert_id
	cert_id := c.GetString("cert_id")
	//到区块链上查询区块数据
	block := blockchain.CHAIN.QueryBlockByCertId(cert_id)
	if block == nil {//遍历了整个区块链但是没有查到数据
		c.Ctx.WriteString("没有查到数据")
		return
	}
	fmt.Println("查询到的区块高度:",block.Height)
	//反序列化
	certRecord,_:= models.DeSerializeCertRecord(block.Data)
	certRecord.CertIdHex= strings.ToUpper(string(certRecord.CertId))
	certRecord.CertHashHex= string(certRecord.CertHash)
	certRecord.CertTimeFormat = tools.TimeFormat(certRecord.CertTime,tools.Time1)
	//结构体
	c.Data["CertRecord"] = certRecord
	//跳转证书详情页面
	c.TplName = "querydata.html"
}