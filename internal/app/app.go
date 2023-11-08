package app

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    // подключаем handler, uscase, repository
    "messanger/internal/usecase"
    "messanger/internal/repository"
    "messanger/internal/handler"
)

func Run() {


    db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")//use config
    repos := repository.NewRepository(db)
    usecases := usecase.NewUseCase(repos)
    handlers := handler.NewHandler(usecases)
    handler.NewRouter(handlers)
}