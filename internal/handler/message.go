package handler

import (
    "fmt"
    "net/http"
    "encoding/json"
    "strconv"
    "log"
    "strings"
)



func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
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

    r.ParseForm()
    message := r.Form.Get("message")
    id, _ := h.uc.Message.Create(message)
    fmt.Fprintf(w, strconv.Itoa(id))
}



func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    params := strings.Split(r.URL.Path, "/")
    id, err := strconv.Atoi(params[len(params)-1])

    res := "{\"status\": \"success\"}"
    if err != nil {
        log.Fatal(err)
        res = "{\"status\": \"error\"}"
    }
    r.ParseForm()
    message := r.Form.Get("message")

    id, _ = h.uc.Message.Update(id, message)
    fmt.Println(id) 
    fmt.Println(id) 
    fmt.Fprintf(w, res) 
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    params := strings.Split(r.URL.Path, "/")

    id, err := strconv.Atoi(params[len(params)-1])

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