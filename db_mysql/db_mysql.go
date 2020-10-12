package db_mysql

import (
	"Datarenzheng1010/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB
func Connect() {
	fmt.Println("链接mysql数据库")
	config := beego.AppConfig //定义config变量,接收并赋值为全局配置变量
	//获取配置选项
	/*appName := config.String("appname")
	fmt.Println("项目应用名称",appName)
	port,err := config.Int("httpport")
	if err != nil {
		panic("项目配置信息解析错误")
	}
	fmt.Println("应用的监听端口",port)*/
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")
	fmt.Println(dbDriver, dbUser, dbPassword)

	//连接数据库
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIp + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil { // err不为nil，表示连接数据库时出现了错误, 程序就在此中断就可以，不用再执行了。
		//早解决，早解决
		panic("数据库连接错误，请检查配置")
	}
	fmt.Println(db)
	Db = db
	//代码封装:可以将重复的代码或者功能相对比较独立的代码，进行封装，以
	//函数的形式进行封装，变成一个代码块或者是功能包，供使用者进行调用

}

func AddUser(u models.User)(int64, error){
	//1、将密码进行hash计算，得到密码hash值，然后在存
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Pwd))
	psswordBytes := md5Hash.Sum(nil)
	u.Pwd = hex.EncodeToString(psswordBytes)
	//execute， .exe可执行文件
	result, err :=Db.Exec("insert into userdata(phone,pwd)" +
		" values(?,?) ", u.Phone,u.Pwd)
	if err != nil {
		return -1,err
	}
	row,err := result.RowsAffected()//用于返回影响数据中几行数据.比如保存了一条数据则
	//返回1
	if err != nil {
		return -1,err
	}
	return row,nil

}