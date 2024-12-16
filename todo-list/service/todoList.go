package service

import (
	"gorm.io/gorm"
	"mygo/todo-list/model"
)

type TodoList struct {
	*gorm.DB
}

func (l *TodoList) GetList() ([]model.TodoList, error) {
	var todos []model.TodoList

	result := l.Order("id desc").Find(&todos)

	if result.Error != nil {
		return nil, result.Error
	}

	return todos, nil
}
