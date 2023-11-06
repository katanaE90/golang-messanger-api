package usecase

import (
	"messanger/entity"
	// "messanger/repository"
	// "messanger/usecase"
)


type Message interface {
	GetAll() ([]entity.Message, error)
	Create() (int, error)
	Update(id int, message string) (int, error)
	Delete(id int) (error)
}

type MessageUseCase struct {
	Message
}



func NewMessageUseCase(repo Message) *MessageUseCase {
	return &MessageUseCase{Message: repo}
}