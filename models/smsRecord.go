package models

import (
	"Datarenzheng1010/db_mysql"
	"fmt"
)

type SmsRecord struct {
	BizId string
	Phone string
	Code string
	Status string
	Message string
	TimeStamp int64
}

//向数据库当中保存验证码记录,该记录由阿里云第三方返回
func (s SmsRecord) SaveSmsRecord()(int64, error){
	rs, err := db_mysql.Db.Exec("insert into sms_record(biz_id,phone,code,status,message,timestamp) values(?,?,?,?,?,?) ",s.BizId,s.Phone,s.Code,s.Status,s.Message,s.TimeStamp)
	if err != nil {
		//保存数据出错
		fmt.Println("数据保存出错!!!",err.Error())
		return -1, err
	}
	return rs.RowsAffected()
}

//根据用户提交的手机号和短信验证码查询验证码是否正确
func QuerySmsRecord(bizId string,phone string,code string)(*SmsRecord,error){
	var sms SmsRecord
	row := db_mysql.Db.QueryRow("select biz_id,timestamp from sms_record where biz_id = ? and phone = ? and code = ?",bizId,phone,code)
	err := row.Scan(&sms.BizId,&sms.TimeStamp)
	if err != nil {
		return nil,err
	}
	return &sms,nil
}