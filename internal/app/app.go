package main

import (
    // "reflect"
    // "fmt"
    // "html"
    "log"
    "net/http"
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


    db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")
    repos := repository.NewRepository(db)
    usecases := usecase.NewUseCase(repos)
    handlers := handler.NewHandler(usecases)

    // httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
    // fmt.Println(handlers)

    // fmt.Printf("OK: h=%#v\n", handlers)
    // fmt.Println(reflect.TypeOf(handlers.MessageHandler))
    

    http.HandleFunc("/messages", handlers.GetAll)
    // fmt.Println(handler.SayHi())

    log.Fatal(http.ListenAndServe(":8180", nil))
}


func main() {
    Run()
}

