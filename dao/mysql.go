package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitMysql() (err error) {
	// 创建数据库
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/dbtest2?charset=utf8mb4&parseTime=True&loc=Local")
	
	// 连接
	if err != nil {
		return
	}
	// 测试是否能连接ping通
	err = DB.DB().Ping()
	return
}
