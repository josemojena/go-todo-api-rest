package test

import (
	"testing"
	"todo-app/Application/Controllers"
	"todo-app/Application/Queries"
	"todo-app/Infraestructure/database"
	"todo-app/Infraestructure/repository"
	"todo-app/Infraestructure/service"
)
var Uow  repository.UnitOfWork
var taskHandler service.TaskCommandHandler
var taskQueries Queries.TaskQueries
var taskController *Controllers.TaskController

func initialize(){

	Uow = repository.InitializeUnitOfWord(&database.DatabaseContext{})
	taskHandler= service.InitializeTaskCommandHandler(Uow)
	taskQueries = Queries.InitializeTaskQueries()
	taskController = Controllers.InitializeTaskController(taskQueries, taskHandler)
}
func Test_AddTask(t *testing.T){
	initialize()
	err:= taskController.AddTask()
	if nil != err {
		t.Errorf("Getting %v when expect nil", err )
	}
}
func Test_GetTasks(t *testing.T) {
  initialize()
  data := taskController.Tasks()
  if len(data) == 0 {
    t.Errorf("No items on the result, got error")
  }

}