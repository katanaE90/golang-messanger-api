package repository

import (
	"messanger/usecase"
)

type Repository struct {
	// Authorization
	Message usecase.Mesage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		// Authorization: NewAuthPostgres(db),
		Message: NewMessageMysql(db),
	}
}
