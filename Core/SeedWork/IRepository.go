package SeedWork

type IRepository interface {
	AddTask()
	UpdateTask()
	RemoveTask()
	ChangeStatus(status bool)
}
