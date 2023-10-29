package repository

// type Authorization interface {
// 	CreateUser(user todo.User) (int, error)
// 	GetUser(username, password string) (todo.User, error)
// }

type Message interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	// GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}



type Repository struct {
	// Authorization
	Message
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		// Authorization: NewAuthPostgres(db),
		Message: NewMessageMysql(db),
	}
}
