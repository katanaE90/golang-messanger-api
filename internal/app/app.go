package main

import (
    "fmt"
_    "html"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    // подключаем handler, uscase, repository

// CRUD
// 1) get messages for chat (limit 50) "SELECT * FROM messages WHERE chat_id = 1"

// 2) delete message "DELETE message FROM messages WHERE id = 1"

// 3) updete message "UPDATE message FROM messages WHERE id = 1"

// 





func main() {
    Run()

}

func Run() {

    // messageUseCase := usecase.New(
    //     repo.New(pg),
    //     webapi.New(),
    // )

    // httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

    http.HandleFunc("/hi", HiHandler)

    log.Fatal(http.ListenAndServe(":8180", nil))
}