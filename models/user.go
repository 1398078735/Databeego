package models

import (
	"Datarenzheng1010/db_mysql"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Pwd string `form:"pwd"`
}

//保存用户方法
func (u User) AddUser() (int64,error) {
	fmt.Println("保存用户信息:",u.Phone,u.Pwd)
	//1、将密码进行hash计算，得到密码hash值，然后在存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Pwd))
	psswordBytes := md5Hash.Sum(nil)
	u.Pwd = hex.EncodeToString(psswordBytes)
	//execute， .exe可执行文件
	rs,err:=db_mysql.Db.Exec("insert into userdata(phone,pwd)" +
		" values(?,?) ", u.Phone,u.Pwd)
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
func (u User) QueryUser()(*User,error){
	fmt.Println("查询用户信息:",u.Phone,u.Pwd)
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Pwd))
	psswordBytes := md5Hash.Sum(nil)
	u.Pwd = hex.EncodeToString(psswordBytes)

	row:=db_mysql.Db.QueryRow("select phone from userdata where phone = ? and pwd = ?",
		u.Phone,u.Pwd)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,err
}

func (u User) QueryUserByphone()(*User,error){
	row:=db_mysql.Db.QueryRow("select id from userdata where phone = ? ",
		u.Phone)

	err := row.Scan(&u.Phone)
	if err != nil {
		return nil,err
	}
	return &u,err
}