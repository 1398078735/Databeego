package controllers

import (
	"Datarenzheng1010/models"
	"Datarenzheng1010/tools"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type SendSmsController struct {
	beego.Controller
}

func (s *SendSmsController) Post(){
	fmt.Println("发送验证码")
	var smsLogin models.SmsLogin
	err := s.ParseForm(&smsLogin)
	if err != nil {
		s.Ctx.WriteString("发送验证码数据解析失败")
		return
	}
	phone := smsLogin.Phone
	code := tools.GenRandCode(6)//返回一个6位的随机数
	result,err:=tools.SendSms(phone,code,tools.SMS_TLP_REGISTER)
	if err != nil {
		s.Ctx.WriteString("发送验证码失败")
		return
	}
	if len(result.BizId) == 0 {
		s.Ctx.WriteString(result.Message)
		return
	}
	//验证码发送成功
	smsRecord := models.SmsRecord{
		BizId:     result.BizId,
		Phone:     phone,
		Code:      code,
		Status:    result.Code,
		Message:   result.Message,
		TimeStamp: time.Now().Unix(),
	}
	_, err = smsRecord.SaveSmsRecord()
	if err!= nil {
		fmt.Println("保存失败",err.Error())
		s.Ctx.WriteString("获取验证码失败")
		return
	}
	//保存成功，bizid
	s.Data["Phone"] = smsLogin.Phone
	s.Data["BizId"] = smsRecord.BizId
	///验证码登录
	s.TplName = "login_sms_second.html"


}