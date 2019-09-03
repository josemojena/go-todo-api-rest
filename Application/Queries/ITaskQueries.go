package Queries

import "todo-app/Application/Dtos"

type ITaskQueries interface {
	GetTasks() ([]Dtos.TaskViewModel, error)
}

