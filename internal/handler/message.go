package handler

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
    "log"
    // "database/sql"
    // _ "github.com/go-sql-driver/mysql"
)



func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        // business as usual
        // h.fs.ServeHTTP(w, r)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    messages, _ := h.uc.Message.GetAll()
    jsonMessages, _ := json.Marshal(messages)
    fmt.Fprintf(w, string(jsonMessages))
}



func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Content-Type", "text/html;charset=utf-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    id, _ := h.uc.Message.Create()
    // jsonMessages, _ := json.Marshal(messages)
    fmt.Fprintf(w, strconv.Itoa(id))
}



func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    id, err := strconv.Atoi(r.URL.Query().Get("id"))

    res := "{\"status\": \"success\"}"
    if err != nil {
        log.Fatal(err)
        res = "{\"status\": \"error\"}"
    }
    message := r.URL.Query().Get("message")

    id, _ = h.uc.Message.Update(id, message)
    fmt.Println(id) 
    fmt.Fprintf(w, res) 
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    fmt.Println(id)
    res := "{\"status\": \"success\"}"

    if err != nil {
        log.Fatal(err)
        res = "{\"status\": \"error\"}"
    }

    err = h.uc.Message.Delete(id)

    if err != nil {
        log.Fatal(err)
        res = "{\"status\": \"error\"}"
    }

    fmt.Fprintf(w, res)        

}