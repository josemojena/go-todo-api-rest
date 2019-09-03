package Queries

import (
	"log"
	"todo-app/Application/Dtos"
	"todo-app/Infraestructure/database"
)

type TaskQueries struct{}

func InitializeTaskQueries() TaskQueries {
	return TaskQueries{}
}
func (tq TaskQueries) GetTasks() ([]Dtos.TaskViewModel, error) {

	tasksviews :=make([]Dtos.TaskViewModel, 0)
    db, err := database.OpenConnection()
    if err != nil {
       return tasksviews, err
	}
    defer db.Close()

    rows, err := db.Query("SELECT id, title, done FROM `task`")
    if err != nil {
    	log.Fatalln(err)
	}
    defer rows.Close()

    for rows.Next() {
		var obj = new(Dtos.TaskViewModel)
		err := rows.Scan(&obj.Id,&obj.Title, &obj.Done)

		if  err == nil {
			tasksviews = append(tasksviews, *obj)
		}

	}
	return tasksviews, nil

}
