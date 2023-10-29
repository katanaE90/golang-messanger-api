package usecase



type MessageUseCase struct {
	repo repository.Message
}

func NewMessageUseCase(repo repository.Message) *MessageUseCase {
	return &MessageUseCase{repo: repo}
}