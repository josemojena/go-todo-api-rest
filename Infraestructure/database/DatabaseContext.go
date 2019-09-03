package database

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"todo-app/Core/Domain"
)

type DatabaseContext struct {
	sqlDb *sql.DB
}
//InitDbConnection a new connection with the database
func InitDbConnection() * DatabaseContext {
	db, err:= OpenConnection()
	if err != nil {
		log.Fatalf("Unable connect to the database, error: %v", err)
	}
	return &DatabaseContext{sqlDb: db}
}

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "wordpress:wordpress@tcp(192.168.99.100:3306)/wordpress")
	return db, err
}
//FindById return a task given it's id
func (db * DatabaseContext) FindById(id string) * Domain.Task {
	var task = new(Domain.Task)
	db.sqlDb.QueryRow("SELECT id, title, done FROM `task` WHERE id=?", id).Scan(&task)
	return task
}
//Insert new task into database
func (db *DatabaseContext) InsertTask(task Domain.Task) error {
	var stmt, err = db.sqlDb.Prepare("INSERT INTO `task` (id, title, done) VALUES (? ,? , ?)")
	if err != nil {
		log.Fatalln("error while preparing statement: %v", err)
	}

	defer stmt.Close()
	result, err := stmt.Exec(task.Id, task.Title, task.Done)

	if err != nil {
		return err
	}
	count, err := result.RowsAffected();
	if err == nil && count > 0  {
		return nil
	}
	return errors.New("Nothing was inserted")
}
