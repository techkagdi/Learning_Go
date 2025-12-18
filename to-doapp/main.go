package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func todoHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, todos) // List all
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	var req struct{ Title string }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todos.Add(req.Title)
	SaveTodos(todos) // Persist
	render.JSON(w, r, map[string]string{"message": "Added!"})
}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	if todos.Complete(id) {
		SaveTodos(todos)
		render.JSON(w, r, map[string]string{"message": "Completed!"})
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	if todos.Delete(id) {
		SaveTodos(todos)
		render.JSON(w, r, map[string]string{"message": "Deleted!"})
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

func main() {
	todos, _ = LoadTodos() // Load on startup

	r := chi.NewRouter()
	r.Get("/todos", todoHandler)
	r.Post("/todos", addTodoHandler)
	r.Put("/todos/{id}", completeTodoHandler) // PUT for complete
	r.Delete("/todos/{id}", deleteTodoHandler)

	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", r)
}
