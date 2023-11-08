package usecase

import (
	"messanger/internal/repository"
)

// duplicate code from repository


type MessageUseCase struct {
	repository.Message
}



func NewMessageUseCase(repo repository.Message) *MessageUseCase {
	return &MessageUseCase{Message: repo}
}