package usecase

import (
	"messanger/repository"
)

type MessageUseCase struct {
	repo Mesage
}

func NewMessageUseCase(repo repository.Message) *MessageUseCase {
	return &MessageUseCase{repo: repo}
}