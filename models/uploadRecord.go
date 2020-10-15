package models

import (
	"Datarenzheng1010/db_mysql"
	"fmt"
)

//上传文件的记录
type UploadRecord struct {
	Id int
	UserId int
	FileName string
	FileSize int
	FileCert string
	FileTitle string
	CertTime int
}

func (u UploadRecord) SavaRecord() (int64,error){
	rs,err:=db_mysql.Db.Exec("insert into upload_record(user_id,file_name,file_size,file_cert,file_title,cert_time)" +
		" values(?,?,?,?,?,?) ", u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)
	if err != nil {
		//保存数据出错
		return -1,err
	}
	_,err = rs.RowsAffected()
	if err != nil {
		return -1,err
	}
	return -1,err
}

//查询用户信息
func (u UploadRecord) QueryUserRecord() (*UploadRecord,error){
	fmt.Println("查询用户信息:",u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)

	row:=db_mysql.Db.QueryRow("select user_id from upload_record where user_id = ? and file_name = ? and file_size= ? and file_cert = ? and file_title = ? and cert_time = ? ",
		u.UserId,u.FileName,u.FileSize,u.FileCert,u.FileTitle,u.CertTime)

	err := row.Scan(&u.UserId)
	if err != nil {
		return nil,err
	}
	return &u,err
}