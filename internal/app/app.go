package main

import (
    // "fmt"
_    "html"
    "log"
    "net/http"
    // "database/sql"
    _ "github.com/go-sql-driver/mysql"

    // подключаем handler, uscase, repository
    "messanger/handler"
    "messanger/repository"
    "messanger/usecase"
)
// CRUD
// 1) get messages for chat (limit 50) "SELECT * FROM messages WHERE chat_id = 1"

// 2) delete message "DELETE message FROM messages WHERE id = 1"

// 3) updete message "UPDATE message FROM messages WHERE id = 1"

// 


func Run() {


    db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")
    repos := repository.NewRepository(db)
    usecases := usecase.NewUseCase(repos)
    handlers := handler.NewHandler(usecase)

    // httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

    http.HandleFunc("/hi", handlers.MessageHandler)
    // fmt.Println(handler.SayHi())

    log.Fatal(http.ListenAndServe(":8180", nil))
}


func main() {
    Run()

}

