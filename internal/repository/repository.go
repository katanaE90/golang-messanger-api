package repository

import (
	"messanger/usecase"
    "database/sql"
)

type Repository struct {
	// Authorization
	usecase.Message
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		// Authorization: NewAuthPostgres(db),
		Message: NewMessageMysql(db),
	}
}
