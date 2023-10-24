package handler

import (
    "fmt"
    "net/http"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    fmt.Fprintf(w, "{\"test\": \"test\"}")
        
    db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")
    res, _ := db.Query("SELECT * FROM messages")

    for res.Next(){
        var message string
        var id int
        err := res.Scan(&id, &message)

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%v\n", message)
    }
    db.Close()
}

// func SayHi() string {
//     return "hello"
// }
