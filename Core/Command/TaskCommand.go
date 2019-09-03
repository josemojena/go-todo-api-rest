package Command
import "github.com/google/uuid"

type TaskCommand struct {
	Id string
	Title string
	Done bool
}
func (tc * TaskCommand) getId() string {
	return tc.Id
}

func CreateTaskCommand(title string, done bool) *TaskCommand {
	return &TaskCommand{ Id : uuid.New().String() , Title:title, Done: done  }
}