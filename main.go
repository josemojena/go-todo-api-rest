package main

import (
	"fmt"
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
var mySqlConnector *database.DatabaseContext
func main(){

	mySqlConnector = database.InitDbConnection()
	Uow = repository.InitializeUnitOfWord(mySqlConnector)
	taskHandler= service.InitializeTaskCommandHandler(Uow)
	taskQueries = Queries.InitializeTaskQueries()
	taskController = Controllers.InitializeTaskController(taskQueries, taskHandler)
	//taskController.AddTask()

	data:= taskController.Tasks()
	fmt.Println(data)
}
