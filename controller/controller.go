package controller

import (
	"github.com/gin-gonic/gin"
	"go/src/models"
	"net/http"
)

func IndexHandler (cont *gin.Context) {
	cont.HTML(http.StatusOK, "index.html", nil)
}

func AddItem (context *gin.Context) {
	//1.从前端的请求中取出数据
	var todo models.Todo
	context.BindJSON(&todo)
	//2.存入数据库
	err := models.AddItem(&todo)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK, todo)
	}
}

func GetItem(context *gin.Context) {
	todolist, err := models.GetALLItem()
	if err != nil{
		context.JSON(http.StatusOK, todolist)
	}else{
		context.JSON(http.StatusOK, todolist)
	}
}

func UpdateItem(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok{
		context.JSON(http.StatusOK, gin.H{"error":"ID不存在"})
	}

	todo, err := models.GetAItem(id)
	if err != nil{
		context.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}
	context.BindJSON(&todo)
	if err = models.UpdateItem(todo); err != nil{
		context.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK, todo)
	}
}

func DeleteItem(context *gin.Context) {
	id, ok := context.Params.Get("id")
	if !ok{
		context.JSON(http.StatusOK, gin.H{"error":"ID不存在"})
		return
	}
	if err := models.DeleteItem(id); err!=nil{
		context.JSON(http.StatusOK, gin.H{"error":err.Error()})
	}else{
		context.JSON(http.StatusOK, gin.H{id:"已删除"})
	}
}
