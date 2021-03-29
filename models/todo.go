package models

import (
	"go/src/dao"
)

// Todo Model
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

//对Model的增删改查操作都在此处进行

func AddItem(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	return
}

func GetALLItem()(todolist []*Todo, err error){
	if err := dao.DB.Find(&todolist).Error; err != nil{
		return nil, err
	}
	return
}

func GetAItem(id string)(todo *Todo, err error){
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err!=nil{
		return nil, err
	}
	return
}

func UpdateItem(todo *Todo)(err error){
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteItem(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}



