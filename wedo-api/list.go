package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WeDo - Todo list with associated Items and collaborators
type WeDo struct {
	ID    int    `json:"ID"`
	Tasks []Task `json:"Tasks"`
}

// CreateList - Creates a Wedo List
func CreateList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am creating a list")
	wedo := WeDo{
		ID: 1,
		Tasks: []Task{
			Task{
				ID:   1,
				Name: "This is wedo item 1",
			},
		},
	}

	j, err := json.Marshal(wedo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// ReadList - Reads a Wedo List
func (l *WeDo) ReadList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am reading a list")
}

// UpdateList - Updates a Wedo List
func (l *WeDo) UpdateList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am updating a list")
}

// DeleteList - Deletes a Wedo List
func (l *WeDo) DeleteList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am deleting a list")
}
