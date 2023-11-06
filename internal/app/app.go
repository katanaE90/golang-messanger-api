package main

import (
    // "reflect"
    // "fmt"
    // "html"
    // "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    // подключаем handler, uscase, repository
    "messanger/repository"
    "messanger/usecase"
    "messanger/handler"
)
// CRUD
// 1) get messages for chat (limit 50) "SELECT * FROM messages WHERE chat_id = 1"

// 2) delete message "DELETE message FROM messages WHERE id = 1"

// 3) updete message "UPDATE message FROM messages WHERE id = 1"

// 


func Run() {


    db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")//use config
    repos := repository.NewRepository(db)
    usecases := usecase.NewUseCase(repos)
    handlers := handler.NewHandler(usecases)
    handler.NewRouter(handlers)



    // httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
    // fmt.Println(handlers)

    // fmt.Printf("OK: h=%#v\n", handlers)
    // fmt.Println(reflect.TypeOf(handlers.MessageHandler))
    


}


func main() {
    Run()
}

