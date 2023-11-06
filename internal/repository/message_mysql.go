package repository

import (
	"messanger/entity"
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
    "context"
)

type MessageMysql struct {
	db *sql.DB
}

func NewMessageMysql(db *sql.DB) *MessageMysql {
	return &MessageMysql{db: db}
}


func (repo *MessageMysql) GetAll() ([]entity.Message, error) {
	var messages []entity.Message

	res, err := repo.db.Query("SELECT * FROM messages")

	for res.Next(){
		var message entity.Message
		err := res.Scan(&message.Id, &message.Text)
		messages = append(messages, message)

        if err != nil {
            log.Fatal(err)
        }
    }

	return messages, err
}



func (repo *MessageMysql) Create() (int, error) {
	var id int

	query := "INSERT INTO `messages` (`message`) VALUES (?)"
	res, err := repo.db.ExecContext(context.Background(),query, "John Doe")

	resId, err := res.LastInsertId()
	id = int(resId)

	if err != nil {
        log.Fatal(err)
    }

	return id, err
}


// что возвращает?
func (repo *MessageMysql) Update(id int, message string) (int, error) {
	query := "UPDATE messages SET message=? WHERE id=?"
	res, err := repo.db.ExecContext(context.Background(),query, message, id)

	resId, err := res.LastInsertId()
	id = int(resId)

	if err != nil {
        log.Fatal(err)
    }

	return id, err
}

func (repo *MessageMysql) Delete(id int) (error) {
	query := "DELETE FROM messages WHERE id=?"
	_, err := repo.db.ExecContext(context.Background(),query, id)

	if err != nil {
        log.Fatal(err)
    }

	return err
}

// func (repo *MessageMysql) Update(id int) (error) {
// 	_, err := repo.db.Query("INSERT INTO messages")

// 	return err
// }