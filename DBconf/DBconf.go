package DBconf

import (
	"database/sql"
	"fmt"
	"github.com/cihub/seelog"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
const (
	USERNAME = "root"
	PASSWORD = "123456"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "mydb"
)



func GetDB() *sql.DB{
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	DB,err :=sql.Open("mysql",dsn)
	//defer DB.Close()

	if err != nil{
		seelog.Error("Open mysql failed,err:%v\n",err)
		return nil
	}
	DB.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)//设置最大连接数
	DB.SetMaxIdleConns(16) //设置闲置连接数

	return DB
}
