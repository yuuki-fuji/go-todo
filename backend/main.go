package main

import (
	"log"
	"net/http"

	"github.com/yuuki-fuji/go-todo/backend/api"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/health", api.Health)
	mux.HandleFunc("/api/todo", api.GetTodos)

	log.Println("server start :8880")
	log.Fatal(http.ListenAndServe(":8880", mux))
}
