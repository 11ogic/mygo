package service

import (
    "gorm.io/gorm"
    "mygo/todo-list/model"
)

type TodoList struct {
    DB *gorm.DB
}

func (l *TodoList) GetList() []model.TodoList {
    var todos []model.TodoList

    l.DB.Order("id desc").Find(&todos)

    return todos
}
