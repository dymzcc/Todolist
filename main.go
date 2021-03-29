package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go/src/controller"
	"go/src/dao"
	"go/src/models"
)



func main()  {
	//创建数据库
	err := dao.InitMySQL()
	if err != nil{
		panic(err)
	}
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})


	engine := gin.Default()
	//静态文件的位置
	engine.Static("/static", "static")
	//HTML文件的位置
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", controller.IndexHandler)


	v1Group := engine.Group("v1")
	{
		//添加
		v1Group.POST("/todo", controller.AddItem)

		//查看所有待办事项
		v1Group.GET("/todo", controller.GetItem)

		//更新,修改某个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateItem)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteItem)
	}

	engine.Run()

}

