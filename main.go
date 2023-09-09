package main

import (
    "fmt"
_    "html"
    "log"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// CRUD 

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
//        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")

 	

        mapD := map[string]int{"apple": 5, "lettuce": 7}
        mapB, _ := json.Marshal(mapD)
        fmt.Println(string(mapB))

//        fmt.Fprintf(w, "{\"test\":\"test\"}")
//        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
//        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")

        fmt.Fprintf(w, "{\"test\":\"test\"}")
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

    })

    log.Fatal(http.ListenAndServe(":8180", nil))

}
