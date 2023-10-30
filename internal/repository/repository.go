package repository

import (
	// "messanger/usecase"
	"messanger/entity"
    "database/sql"
)

type Message interface {
	GetAll() ([]entity.Message, error)
}


type Repository struct {
	// Authorization
	Message
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		// Authorization: NewAuthPostgres(db),
		Message: NewMessageMysql(db),
	}
}
