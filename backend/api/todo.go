package api

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]any{
		"items": []Todo{
			{ID: "1", Title: "最初のTodo", Status: "completed"},
			{ID: "2", Title: "２個目のTodo", Status: "in-progress"},
		},
	}

	_ = json.NewEncoder(w).Encode(resp)
}
