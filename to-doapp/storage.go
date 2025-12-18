package main

import (
	"encoding/json"
	"os"
)

const fileName = "todos.json"

func SaveTodos(tl TodoList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}

func LoadTodos() (TodoList, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err // File doesn't exist? Fine, return empty
	}
	var tl TodoList
	err = json.Unmarshal(data, &tl)
	return tl, err
}
