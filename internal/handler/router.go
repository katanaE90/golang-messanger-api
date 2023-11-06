package handler

import (
	"net/http"
    "log"
    // "fmt"
)

func NewRouter(handlers *Handler) {
	// Options
	// fmt.Println(http.Method)
    http.HandleFunc("/messages", handlers.GetAll).Methods("GET")
    // http.HandleFunc("/messages", handlers.Create)
    http.HandleFunc("/messages/create", handlers.Create)
    http.HandleFunc("/messages/update", handlers.Update)
    http.HandleFunc("/messages/delete", handlers.Delete)

    log.Fatal(http.ListenAndServe(":8180", nil))
}