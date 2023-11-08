package repository

import (
	"messanger/internal/entity"
)


type Message interface {
	GetAll() ([]entity.Message, error)
	Create(message string) (int, error)
	Update(id int, message string) (int, error)
	Delete(id int) (error)
}