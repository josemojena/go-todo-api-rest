package SeedWork


type ICommandHandler interface {
	AddTask(command ICommand) error
	UpdateTask(command ICommand) error
	DeleteTask(command ICommand) error
}