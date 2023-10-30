package handler

import (
    "fmt"
    "net/http"
    "encoding/json"
    // "log"
    // "database/sql"
    // _ "github.com/go-sql-driver/mysql"
)

func (h *Handler) MessageHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    fmt.Fprintf(w, "{\"test\": \"test\"}")
        
    // db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")
    // res, _ := db.Query("SELECT * FROM messages")

    // for res.Next(){
    //     var message string
    //     var id int
    //     err := res.Scan(&id, &message)

    //     if err != nil {
    //         log.Fatal(err)
    //     }

    //     fmt.Printf("%v\n", message)
    // }
    // db.Close()
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    messages, _ := h.uc.Message.GetAll()
    jsonMessages, _ := json.Marshal(messages)
    fmt.Fprintf(w, string(jsonMessages))
    
    // fmt.Fprintf("mgs: %#v", a)

    // fmt.Fprintf(w, "{\"test\": \"test\"}")
        
}

// func SayHi() string {
//     return "hello"
// }
