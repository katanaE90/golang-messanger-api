package repository

import (
	"messanger/entity"
	"database/sql"
	"log"
    _ "github.com/go-sql-driver/mysql"
)

type MessageMysql struct {
	db *sql.DB
}

func NewMessageMysql(db *sql.DB) *MessageMysql {
	return &MessageMysql{db: db}
}

// func (r *MessageMysql) Create(userId int, list todo.TodoList) (int, error) {
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return 0, err
// 	}

// 	var id int
// 	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
// 	row := tx.QueryRow(createListQuery, list.Title, list.Description)
// 	if err := row.Scan(&id); err != nil {
// 		tx.Rollback()
// 		return 0, err
// 	}

// 	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
// 	_, err = tx.Exec(createUsersListQuery, userId, id)
// 	if err != nil {
// 		tx.Rollback()
// 		return 0, err
// 	}

// 	return id, tx.Commit()
// }

func (r *MessageMysql) GetAll() ([]entity.Message, error) {
	var messages []entity.Message

    db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")
	res, err := db.Query("SELECT * FROM messages")

	for res.Next(){
		var message entity.Message
		err := res.Scan(&message.Id, &message.Text)
		messages = append(messages, message)
//	messages["id"] = strconv.Itoa(id)
//	messages["text"] = message

        if err != nil {
            log.Fatal(err)
        }
        //        fmt.Printf("id: %v; ", message.id)
        //        fmt.Printf("message: %v\n", message.text)
    }

	return messages, err
}



// func (r *MessageMysql) Delete(userId, listId int) error {
// 	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
// 		todoListsTable, usersListsTable)
// 	_, err := r.db.Exec(query, userId, listId)

// 	return err
// }

// func (r *MessageMysql) Update(userId, listId int, input todo.UpdateListInput) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.Title != nil {
// 		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
// 		args = append(args, *input.Title)
// 		argId++
// 	}

// 	if input.Description != nil {
// 		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
// 		args = append(args, *input.Description)
// 		argId++
// 	}

// 	// title=$1
// 	// description=$1
// 	// title=$1, description=$2
// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
// 		todoListsTable, setQuery, usersListsTable, argId, argId+1)
// 	args = append(args, listId, userId)

// 	logrus.Debugf("updateQuery: %s", query)
// 	logrus.Debugf("args: %s", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }


// func (r *MessageMysql) GetById(userId, listId int) (todo.TodoList, error) {
// 	var list todo.TodoList

// 	query := fmt.Sprintf(`SELECT tl.id, tl.title, tl.description FROM %s tl
// 								INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`,
// 		todoListsTable, usersListsTable)
// 	err := r.db.Get(&list, query, userId, listId)

// 	return list, err
// }