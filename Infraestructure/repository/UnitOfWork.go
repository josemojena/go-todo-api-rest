package repository

import (
	"todo-app/Core/Domain"
	"todo-app/Infraestructure/database"
)

type UnitOfWork struct {
    DatabaseContext * database.DatabaseContext
}
func InitializeUnitOfWord(databaseContext * database.DatabaseContext)  UnitOfWork {
	return UnitOfWork{ DatabaseContext :databaseContext}
}
func (uof * UnitOfWork) AddTask(domain Domain.Task ) error {
	return uof.DatabaseContext.InsertTask(domain)
}
func (UnitOfWork) UpdateTask(domain Domain.Task ){

}
func (UnitOfWork) DeleteTask(domain Domain.Task ){

}

func (uof * UnitOfWork)Find(id string)  Domain.Task {
	return *uof.DatabaseContext.FindById(id)
}