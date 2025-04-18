package models

import (
	"bubble_front/dao"
	"fmt"
)

// 放置所有的模型 ,以及对其进行的crud的操作都放这里 ,只写最基础的方法
// 尽可能的减少业务逻辑

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// Todo crud
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	if err != nil {
		return
	}
	return
}

func GetTodoList() (todoList []*Todo, err error) { // todoList在此定义过了，下面无需定义
	err = dao.DB.Find(&todoList).Error // 这里不需要new了，因为这里对底层数组的引用
	return todoList, err
}

func GetATodo(id string) (todo *Todo, err error) {
	// 如果没有   todo = new(Todo)   这一行，即   todo   是一个   Todo   类型的零值，那么   dao.DB.Where("id=?", id).First(todo)   会报错，因为   gorm   不支持将结果直接填充到一个值类型变量中。
	// 使用指针作为数据库查询结果的目的地是 Go 语言中常见的做法，它允许直接修改原始数据，避免数据复制，并提供更灵活的错误处理方式。因此，在执行数据库查询并将结果填充到变量时，确保使用指针类型是非常重要的。
	todo = new(Todo) // 不理解 ,unsupported destination, should be slice or struct
	err = dao.DB.Debug().Where("id=?", id).First(todo).Error
	return todo, err
}

func UpdateAtodo(todo *Todo) (err error) {
	// if err = dao.DB.Save(todo).Error; err != nil {
	// 	return
	// }
	err = dao.DB.Save(todo).Error
	fmt.Println("todo:", todo)
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
