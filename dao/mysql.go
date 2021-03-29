package dao

import "github.com/jinzhu/gorm"



//声明一个全局变量
var (DB *gorm.DB)

func InitMySQL()(err error){
	dsn := "root:dede@tcp(39.97.212.145)/gindb?charset=utf8mb4&parseTime=True&loc=Local"
	DB , err = gorm.Open("mysql", dsn)
	if err != nil{
		return
	}
	//测试连通性
	err = DB.DB().Ping()
	return
}