package Controllers

import (
	"todo-app/Application/Dtos"
	"todo-app/Application/Queries"
	"todo-app/Core/Command"
	"todo-app/Infraestructure/service"
)

type TaskController struct {
	query   Queries.ITaskQueries
	handler service.TaskCommandHandler
}

func InitializeTaskController(query Queries.ITaskQueries, handler service.TaskCommandHandler) *TaskController {
	return &TaskController{query: query, handler: handler}
}
func (tc *TaskController) Tasks() [] Dtos.TaskViewModel {
	data, err := tc.query.GetTasks()
	if err != nil {
		panic(err)
	}
	return data
}
func (tc *TaskController) AddTask() error {
	command := Command.CreateTaskCommand("hello word", false)
	return tc.handler.CreateTaskCommandHandler(*command)

}
func (tc * TaskController) RemoveTask() {}
