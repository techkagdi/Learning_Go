package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoList []Todo // Slice of Todos – our "database"

var todos TodoList // Global slice (for simplicity; use maps for IDs later)

func (tl *TodoList) Add(title string) {
	id := len(*tl) + 1 // Simple ID gen
	*tl = append(*tl, Todo{
		ID:        id,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	})
	fmt.Printf("Added todo: %s (ID: %d)\n", title, id)
}

func (tl *TodoList) List() {
	if len(*tl) == 0 {
		fmt.Println("No todos yet!")
		return
	}
	for _, todo := range *tl { // for-range loop!
		status := "Pending"
		if todo.Completed {
			status = "Done"
		}
		fmt.Printf("[%d] %s - %s (Created: %s)\n", todo.ID, status, todo.Title, todo.CreatedAt.Format("2006-01-02"))
	}
}

func (tl *TodoList) Complete(id int) bool {
	for i := range *tl { // Loop with index
		if (*tl)[i].ID == id {
			(*tl)[i].Completed = true
			fmt.Printf("Completed todo: %s\n", (*tl)[i].Title)
			return true
		}
	}
	fmt.Println("Todo not found!")
	return false
}

func (tl *TodoList) Delete(id int) bool {
	for i := range *tl {
		if (*tl)[i].ID == id {
			fmt.Printf("Deleted todo: %s\n", (*tl)[i].Title)
			*tl = append((*tl)[:i], (*tl)[i+1:]...) // Slice delete
			return true
		}
	}
	fmt.Println("Todo not found!")
	return false
}
