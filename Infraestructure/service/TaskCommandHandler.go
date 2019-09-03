package service

import (
	"todo-app/Core/Command"
	"todo-app/Core/Domain"
	"todo-app/Infraestructure/repository"
)

type TaskCommandHandler struct {
	uoW  repository.UnitOfWork
}

func InitializeTaskCommandHandler(work repository.UnitOfWork) TaskCommandHandler {
	return TaskCommandHandler{uoW: work}
}
func (tch *TaskCommandHandler) CreateTaskCommandHandler(taskCommand Command.TaskCommand) error {
	task := Domain.Task{
		Id:    taskCommand.Id,
		Title: taskCommand.Title,
		Done:  taskCommand.Done,
	}
	return tch.uoW.AddTask(task)
}
