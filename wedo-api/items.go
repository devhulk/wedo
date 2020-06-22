package api

import (
	"fmt"
	"net/http"
)

// Task - a list item with associated Details
type Task struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

// CreateItem - Creates a Wedo Item
func (l *Task) CreateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am creating a item")
}

// ReadItem - Reads a Wedo Item
func (l *Task) ReadItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am reading a item")
}

// UpdateItem - Updates a Wedo Item
func (l *Task) UpdateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am updating a item")
}

// DeleteItem - Deletes a Wedo Item
func (l *Task) DeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am deleting a item")
}
