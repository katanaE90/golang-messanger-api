package usecase


import (
	"messanger/internal/repository"
)


type UseCase struct {
	// Authorization
	repository.Message
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		// Authorization: NewAuthService(repos.Authorization),
		Message: NewMessageUseCase(repos.Message),
	}
}