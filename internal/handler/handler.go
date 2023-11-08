package handler

import (
	// "fmt"
    // "net/http"
	"messanger/internal/usecase"
)

type Handler struct {
	uc     *usecase.UseCase
	//logger logger.Logger
}

func NewHandler(uc *usecase.UseCase) *Handler {
	return &Handler{uc: uc}
}
