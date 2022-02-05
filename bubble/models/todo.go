package models

import (
	"gin_demo/bubble/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo 增删改查

func CreateTodo(todo *Todo) (err error) {

	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todo *[]Todo, err error) {
	todo = new([]Todo)

	err = dao.DB.Find(todo).Error
	if err != nil {
		return nil, err
	} else {
		return todo, err
	}
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Where("id = ?", id).First(todo).Error
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}
