package controllers

import (
	"Datarenzheng1010/models"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type HomeController struct {
	beego.Controller
}


//post用io包保存文件
//post1用beego框架的方式保存文件
func (h *HomeController) Post(){
	phone :=h.Ctx.Request.PostFormValue("phone")

	//该post的方法用于处理用户在客户端的文件
	title := h.Ctx.Request.PostFormValue("upload_title")//获取用户输入信息
	fmt.Println(title)
	//用户上传的文件
	file,header,err:=h.GetFile("tengyuanqianhua")
	if err != nil {//解析客户端提交的文件出现错误
		h.Ctx.WriteString("抱歉文件解析失败")
		return
	}
	defer file.Close()

	fmt.Println("自定义的标题:",title)
	//获得到上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
	//支持jpg，png类型
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(":",strings.TrimSpace(fileType))
	isJpg:=strings.HasSuffix(header.Filename,".jpg")
	isPng:=strings.HasSuffix(header.Filename,".png")
	//判断文件格式
	if !isJpg && !isPng {
		h.Ctx.WriteString("抱歉格式出错")
		return
	}
	//文件的大小
	config := beego.AppConfig
	fileSize,err:=config.Int64("file_size")
	fmt.Println("上传的文件的大小:",header.Size)//返回字节大小

	if header.Size / 1024 > fileSize {
		h.Ctx.WriteString("抱歉文件大小超出范围")
		return
	}
	//使用os的包提供的方法保存文件
	//int64数据的长度
	savaFilePath := "static/upload" + header.Filename
	saveFile,err:=os.OpenFile(savaFilePath,os.O_CREATE|os.O_RDWR,777)
	if err != nil {
		h.Ctx.WriteString("抱歉,文件保存失败")
		return
	}
	_,err = io.Copy(saveFile,file)
	if err != nil {
		h.Ctx.WriteString("抱歉,电子数据认证失败")
		return
	}

	//计算文件的SMA256值
	hash256:=sha256.New()
	fileBytes,_:=ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes := hash256.Sum(nil)
	fmt.Println(hex.EncodeToString(hashBytes))

	//先查询用户id
	user1,err := models.User{Phone:phone}.QueryUserByphone()
	if err != nil {
		h.Ctx.WriteString("抱歉，查询用户失败")
	}


	//把上传的文件作为记录保存到数据库当中
	md5Has:=md5.New()
	mdhfileBytes,err:=ioutil.ReadAll(saveFile)
	md5Has.Write(mdhfileBytes)
	bytes := md5Has.Sum(nil)
	record := models.UploadRecord{
		UserId:    user1.Id,
		FileName:  header.Filename,
		FileSize:  header.Size,
		FileCert:  hex.EncodeToString(bytes),
		FileTitle: title,
		CertTime:  time.Now().Unix(),
	}
	_ , err = record.SavaRecord()
	if err != nil {
		h.Ctx.WriteString("抱歉数据保存失败")
		return
	}
	//上传文件保存到数据库中
	records,err:=models.QueryUserRecord(user1.Id)
	if err != nil {
		h.Ctx.WriteString("抱歉获取数据失败")
		return
	}

	h.Data["Records"] = records
	h.TplName = "uploadRecord.html"

}


func (h *HomeController) Post1() {
	//用户上传的自定义的标题
	title := h.Ctx.Request.PostFormValue("upload_title")//获取用户输入信息

	//用户上传的文件
	file,header,err:=h.GetFile("tengyuanqianhua")
	if err != nil {//解析客户端提交的文件出现错误
		h.Ctx.WriteString("抱歉文件解析失败")
		return
	}
	defer file.Close()
	fmt.Println("自定义的标题:",title)
	//获得到上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
	//支持jpg，png类型
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(":",strings.TrimSpace(fileType))
	isJpg:=strings.HasSuffix(header.Filename,".jpg")
	isPng:=strings.HasSuffix(header.Filename,".png")
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
	fileSize,err:=config.Int64("file_size")
	fmt.Println("上传的文件的大小:",header.Size)//返回字节大小

	if header.Size / 1024 > fileSize {
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


	_,err = os.Open(saveDir)
	//打开某个文件用openfile;打开某个目录用os.opne
	if err != nil {
		err = os.Mkdir(saveDir,777)
		if err != nil {
			h.Ctx.WriteString("抱歉文件认证遇到错误")
			return
		}
	}

	saveName := saveDir+"/"+header.Filename
	fmt.Println("要保存的文件名",saveName)

	err = h.SaveToFile("tengyuanqianhua",saveName)
	if err != nil {
		h.Ctx.WriteString("抱歉文件认证失败")
		return
	}

	h.Ctx.WriteString("获取到上传文件")

	fmt.Println(file)
}

