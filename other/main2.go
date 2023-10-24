package main

import (
_    "reflect"
    "encoding/json"
_    "strconv"
    "fmt"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


type Message struct {
    Id int `json:"id"`
    Text string `json:"text"`
}

// CRUD 

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")


        messages := map[string]string{"id": "5", "text": "hello", "user_id": "1"}


        jsonMessages, _ := json.Marshal(messages)
//        fmt.Println(reflect.TypeOf(messages))
        fmt.Println(string(jsonMessages))
        fmt.Fprintf(w, string(jsonMessages))

    })

    http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request){
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")

        db, _ := sql.Open("mysql", "root:Qwe123??@tcp(127.0.0.1:3306)/messanger")
        res, _ := db.Query("SELECT * FROM messages")


//        messages := map[string]string{"id": "5", "text": "hello"}
	messages := []Message{}
for res.Next(){
	var message Message
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
        db.Close()

        jsonMessages, _ := json.Marshal(messages)
        fmt.Fprintf(w, string(jsonMessages))
    })

    log.Fatal(http.ListenAndServe(":8180", nil))

}
