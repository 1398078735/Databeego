package tools

import (
	"encoding/json"
	"fmt"
	"github.com/alibaba-cloud-sdk-go-master/services/dysmsapi"
	"github.com/astaxie/beego"
	"math/rand"
	"strings"
	"time"
)

const SMS_TLP_REGISTER  = "SMS_205393604"//注册业务的短信模板
const SMS_TLP_LOGIN =  ""//用户登录的短信模板
const SMS_TLP_KYC  = ""//实名认证的短信模板
type SmsCode struct {
	Code string `json:"code"`
}

type SmsResutl struct {
	BizId string
	Code  string
	Message string
	RequestId string
}


//该函数用于发送一条短信息
//phone电话接受验证码的号码
//code发送的验证码信息
//templtype：模板类型
func SendSms(phone string,code string,templateType string)(*SmsResutl,error){
	config := beego.AppConfig
	accessKey := config.String("sms_access_key")
	accessKeySecret := config.String("sms_access_secret")
	client,err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",accessKey,accessKeySecret)
	if err != nil{
		return nil,err
	}
	//batch:批量
	request := dysmsapi.CreateSendSmsRequest()
	request.PhoneNumbers = phone//指定要发送给的目标手机号
	request.SignName = "线上餐厅"//签名信息
	request.TemplateCode = templateType //指定短信模板
	//{"code":"xxxx"}:json格式
	smsCode := SmsCode{
		Code:code,
	}
	smsbytes,_:=json.Marshal(smsCode)
	request.TemplateParam = string(smsbytes)//指定要发送的验证码

	response,err:=client.SendSms(request)
	if err != nil {
		return nil,err
	}
	//biz:business,商业，业务
	smsResult := &SmsResutl{
		BizId:    response.BizId,
		Code:      response.Code,
		Message:   response.Message,
		RequestId: response.RequestId,
	}
	return smsResult,err
}

//生成一个位数为width的随机验证码,并将该验证码返回
func GenRandCode(width int)string{
	numeric := [10]byte{0,1,2,3,4,5,6,7,8,9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[ rand.Intn(r) ])
	}
	return sb.String()

}