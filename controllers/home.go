package controllers

import (
	"Datarenzheng1010/blockchain"
	"Datarenzheng1010/models"
	"Datarenzheng1010/tools"
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
	"time"
)

type HomeController struct {
	beego.Controller
}

//post用io包保存文件
//post1用beego框架的方式保存文件
func (h *HomeController) Post() {
	phone := h.Ctx.Request.PostFormValue("phone")

	//该post的方法用于处理用户在客户端的文件
	title := h.Ctx.Request.PostFormValue("upload_title") //获取用户输入信息
	//用户上传的文件
	file, header, err := h.GetFile("tengyuanqianhua")
	if err != nil { //解析客户端提交的文件出现错误
		h.Ctx.WriteString("抱歉文件解析失败")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标题:", title)
	//获得到上传的文件
	fmt.Println("上传的文件名称:", header.Filename)
	//支持jpg，png类型
	fileNameSlice := strings.Split(header.Filename, ".")
	fileType := fileNameSlice[1]
	fmt.Println(":", strings.TrimSpace(fileType))
	isJpg := strings.HasSuffix(header.Filename, ".jpg")
	isPng := strings.HasSuffix(header.Filename, ".png")
	//判断文件格式
	if !isJpg && !isPng {
		h.Ctx.WriteString("抱歉格式出错")
		return
	}
	//文件的大小
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")
	fmt.Println("上传的文件的大小:", header.Size) //返回字节大小

	if header.Size/1024 > fileSize {
		h.Ctx.WriteString("抱歉文件大小超出范围")
		return
	}
	//使用os的包提供的方法保存文件
	//int64数据的长度
	savaFilePath := "static/upload" + header.Filename
	_ , err = tools.SaveFile(savaFilePath,file)
	if err != nil {
		h.Ctx.WriteString("抱歉,文件保存失败")
		return
	}
	/*saveFile, err := os.OpenFile(savaFilePath, os.O_CREATE|os.O_RDWR, 777)
	if err != nil {
		h.Ctx.WriteString("抱歉,文件保存失败")
		return
	}
	_, err = io.Copy(saveFile, file)
	if err != nil {
		h.Ctx.WriteString("抱歉,电子数据认证失败")
		return
	}8*/

	//计算文件的SMA256值
	fileHash := tools.SHA256HashReader(file)
	fmt.Println(fileHash)

	//先查询用户id
	fmt.Println("要查询的用户的phone:", phone)
	user1, err := models.User{Phone: phone}.QueryUserByphone()
	fmt.Println("查询到的用户:", user1)
	if err != nil {
		h.Ctx.WriteString("抱歉，查询用户失败")
	}

	//把上传的文件作为记录保存到数据库当中
	saveFile,err := os.Open(savaFilePath)
	md5String := tools.Md5HashReader(saveFile)
	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size/1024,
		FileCert:  md5String,
		FileTitle: title,
		CertTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err = record.SavaRecord()
	if err != nil {
		fmt.Println(err.Error())
		h.Ctx.WriteString("抱歉数据保存失败")
		return
	}

	//将用户上传的文件的md5和sha256值保存到区块链上，即数据上链
	user := &models.User{
		Phone: phone,
	}
	user,_ = user.QueryUserByphone()
	certRecord := models.CertRecord{
		CertId:   []byte(md5String),
		CertHash: []byte(fileHash),
		CertName: user.Name,
		Phone:    user.Phone,
		CertCard: user.Card,
		FileName: header.Filename,
		FileSize: header.Size/1024,
		CertTime: time.Now().Unix(),
	}
	//序列化
	certBytes,_:=certRecord.Serialize()
	block,err := blockchain.CHAIN.SaveData(certBytes)
	if err!= nil {
		fmt.Println(err.Error())
		h.Ctx.WriteString("数据上链失败")
		return
	}
	fmt.Println("恭喜,已经将数据保存到区块链上,区块高度是",block.Height)

	//blocks,_:=blockchain.CHAIN.QueryAllBlocks()
	//for _, block := range blocks{
	//	fmt.Printf("区块高度:%d,区块内数据:%s\n",block.Height,string(block.Data))
	//}


	//上传文件保存到数据库中
	records, err := models.QueryUserRecord(user1.Id)
	if err != nil {
		fmt.Println(err.Error())
		h.Ctx.WriteString("抱歉获取数据失败")
		return
	}

	h.Data["Records"] = records
	h.TplName = "uploadRecord.html"

}

func (h *HomeController) Post1() {
	//用户上传的自定义的标题
	title := h.Ctx.Request.PostFormValue("upload_title") //获取用户输入信息

	//用户上传的文件
	file, header, err := h.GetFile("tengyuanqianhua")
	if err != nil { //解析客户端提交的文件出现错误
		h.Ctx.WriteString("抱歉文件解析失败")
		return
	}
	defer file.Close()
	fmt.Println("自定义的标题:", title)
	//获得到上传的文件
	fmt.Println("上传的文件名称:", header.Filename)
	//支持jpg，png类型
	fileNameSlice := strings.Split(header.Filename, ".")
	fileType := fileNameSlice[1]
	fmt.Println(":", strings.TrimSpace(fileType))
	isJpg := strings.HasSuffix(header.Filename, ".jpg")
	isPng := strings.HasSuffix(header.Filename, ".png")
	//判断文件格式
	if !isJpg && !isPng {
		h.Ctx.WriteString("抱歉格式出错")
		return
	}
	/*if fileType != "jpg" || fileType != "png"{
		//文件类型不支持
		h.Ctx.WriteString("抱歉文件类型不支持")
		return
	}*/

	//文件的大小
	config := beego.AppConfig
	fileSize, err := config.Int64("file_size")
	fmt.Println("上传的文件的大小:", header.Size) //返回字节大小

	if header.Size/1024 > fileSize {
		h.Ctx.WriteString("抱歉文件大小超出范围")
		return
	}

	//fromfile:文件
	//tofile:要保存的文件
	//perm:permission 权限
	//a:文件所有者对文件的操作权限 读4,写2,执行1
	//b:文件所有者所在组的用户的操作权限
	//c：其他用户的操作权限
	//先尝试打开文件夹
	saveDir := "static/upload"

	//os.OpenFile("文件名",os.O_CREATE|os.O_RDONLY,777)

	_, err = os.Open(saveDir)
	//打开某个文件用openfile;打开某个目录用os.opne
	if err != nil {
		err = os.Mkdir(saveDir, 777)
		if err != nil {
			h.Ctx.WriteString("抱歉文件认证遇到错误")
			return
		}
	}

	saveName := saveDir + "/" + header.Filename
	fmt.Println("要保存的文件名", saveName)

	err = h.SaveToFile("tengyuanqianhua", saveName)
	if err != nil {
		h.Ctx.WriteString("抱歉文件认证失败")
		return
	}

	h.Ctx.WriteString("获取到上传文件")

	fmt.Println(file)
}
