package usecase


import (
	"messanger/repository"
)
// type Authorization interface {
// 	CreateUser(user todo.User) (int, error)
// 	GenerateToken(username, password string) (string, error)
// 	ParseToken(token string) (int, error)
// }

// type Message interface {
// 	// Create(userId int, list todo.TodoList) (int, error)
// 	GetAll() ([]entity.Message, error)
// 	// GetById(userId, listId int) (todo.TodoList, error)
// 	// Delete(userId, listId int) error
// 	// Update(userId, listId int, input todo.UpdateListInput) error
// }


// type Repository struct {
// 	// Authorization
// 	Message
// }



type UseCase struct {
	// Authorization
	Message
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		// Authorization: NewAuthService(repos.Authorization),
		Message: NewMessageUseCase(repos.Message),
	}
}