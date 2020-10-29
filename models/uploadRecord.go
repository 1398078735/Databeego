package models

import (
	"Datarenzheng1010/db_mysql"
	"fmt"
)

//上传文件的记录
type UploadRecord struct {
	Id        int
	UserId    int
	FileName  string
	FileSize  int64
	FileCert  string
	FileTitle string
	CertTime  string
}

//把一条认证数据保存到数据库表中
func (u UploadRecord) SavaRecord() (int64, error) {
	fmt.Println("要保存的认证数据对应的用户的id：", u.UserId)
	rs, err := db_mysql.Db.Exec("insert into upload_record(user_id,file_name,file_size,file_cert,file_title,cert_time)"+
		" values(?,?,?,?,?,?) ", u.UserId, u.FileName, u.FileSize, u.FileCert, u.FileTitle, u.CertTime)
	if err != nil {
		//保存数据出错
		return -1, err
	}
	_, err = rs.RowsAffected()
	if err != nil {
		return -1, err
	}
	return -1, err
}

//查询用户信息
func QueryUserRecord(userId int) ([]UploadRecord, error) {
	rs, err := db_mysql.Db.Query("select * from upload_record where user_id = ?",userId)
	if err != nil {
		return nil, err
	}
	//从rs中读取查询到的数据,返回
	Records := make([]UploadRecord, 0) //容器
	for rs.Next() {
		var record UploadRecord
		err := rs.Scan(&record.Id,&record.UserId, &record.FileName, &record.FileSize, &record.FileCert, &record.FileTitle, &record.CertTime)
		if err != nil {
			return nil, err
		}
		Records = append(Records, record)
	}
	return Records, nil
}
