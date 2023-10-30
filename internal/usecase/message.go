package usecase

import (
	"messanger/entity"
	// "messanger/repository"
	// "messanger/usecase"
)

type MessageUseCase struct {
	repo entity.Message
}

func NewMessageUseCase(repo usecase.Message) *MessageUseCase {
	return &MessageUseCase{repo: repo}
}