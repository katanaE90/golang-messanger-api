package handler

import (
	"messanger/usecase"
)

type Handler struct {
	uc     usecase.UseCase
	//logger logger.Logger
}

func NewHandler(uc UseCase) *Handler {
	return &Handler{uc: uc}
}