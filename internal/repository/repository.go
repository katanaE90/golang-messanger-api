package repository

import (
	// "messanger/usecase"
    "database/sql"
)


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
