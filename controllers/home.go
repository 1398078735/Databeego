package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type HomeController struct {
	beego.Controller
}

func (h *HomeController) Post() {
	//用户上传的自定义的标题
	title := h.Ctx.Request.PostFormValue("upload_title")//从用户输入信息

	//用户上传的文件
	file,header,err:=h.GetFile("tengyuanqianhua")
	if err != nil {//解析客户端提交的文件出现错误
		h.Ctx.WriteString("抱歉文件解析失败")
		return
	}
	fmt.Println("自定义的标题:",title)
	//获得到上传的文件
	fmt.Println("上传的文件名称:",header.Filename)
	//支持jpg，png类型
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	fmt.Println(":",fileType)
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

	if header.Size / 1024 > fileSize {
		h.Ctx.WriteString("抱歉文件大小超出范围")
		return
	}

	fmt.Println("上传的文件的大小:",header.Size)//返回字节大小
	h.Ctx.WriteString("获取到上传文件")

	fmt.Println(file)
}

